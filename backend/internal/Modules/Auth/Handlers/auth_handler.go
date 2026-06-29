package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"fmt"
	"net/http"
	"sort"
	"strings"

	mw       "school-ms/internal/middleware"
	dtos     "school-ms/internal/Modules/Auth/DTOs"
	services "school-ms/internal/Modules/Auth/Services"
	"school-ms/internal/pkg/response"
)

type AuthHandler struct{ svc *services.AuthService }

func NewAuthHandler(svc *services.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// POST /auth/login
// POST /auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    tenant := mw.GetResolvedTenant(r.Context())

    slog.Info("login request received",
        "method", r.Method,
        "path", r.URL.Path,
        "remote_addr", r.RemoteAddr,
        "tenant_id", tenant.ID,
        "host", r.Host,
    )

    if tenant.ID == 0 {
        slog.Warn("tenant resolution failed",
            "host", r.Host,
            "forwarded_host", r.Header.Get("X-Forwarded-Host"),
        )

        response.Unauthorized(
            w,
            "tenant could not be resolved — check Host header or X-Forwarded-Host",
        )
        return
    }

    defer r.Body.Close()

    var req dtos.LoginRequest

    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    if err := dec.Decode(&req); err != nil {
        slog.Warn("login JSON decode failed",
            "tenant_id", tenant.ID,
            "error", err.Error(),
        )

        response.BadRequest(
            w,
            "invalid JSON: "+sanitiseDecodeError(err),
        )
        return
    }

    req.Normalize()

    slog.Info("login attempt",
        "tenant_id", tenant.ID,
        "email", req.Email,
    )

    if errs := req.Validate(); errs != nil {
        slog.Warn("login validation failed",
            "tenant_id", tenant.ID,
            "email", req.Email,
            "errors", errs,
        )

        response.BadRequest(
            w,
            formatValidationErrors(errs),
        )
        return
    }

    res, err := h.svc.Login(req, tenant.ID)

    if err != nil {
        slog.Warn("login failed",
            "tenant_id", tenant.ID,
            "email", req.Email,
            "error", err.Error(),
        )

        switch {
        case errors.Is(err, services.ErrAccountDisabled):
            response.Forbidden(
                w,
                "this account has been disabled — contact your administrator",
            )
        default:
            response.Unauthorized(
                w,
                "invalid email or password",
            )
        }

        return
    }

    slog.Info("login successful",
        "tenant_id", tenant.ID,
        "email", req.Email,
        "user_id", res.User.ID, // adjust field if different
    )

    response.Success(w, res, "login successful")
}

// POST /auth/refresh
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dtos.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "invalid JSON body")
		return
	}
	if strings.TrimSpace(req.RefreshToken) == "" {
		response.BadRequest(w, "refresh_token is required")
		return
	}
	res, err := h.svc.Refresh(req.RefreshToken)
	if err != nil {
		response.Unauthorized(w, "invalid or expired refresh token")
		return
	}
	response.Success(w, res, "token refreshed")
}

// GET /auth/me
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp := map[string]interface{}{
		"user_id":   mw.GetUserID(ctx),
		"tenant_id": mw.GetTenantID(ctx),
		"school_id": mw.GetSchoolID(ctx),
		"roles":     mw.GetRoles(ctx),
	}
	if ayID := mw.GetAcademicYearID(ctx); ayID != nil {
		resp["academic_year_id"] = *ayID
	}
	if tID := mw.GetTermID(ctx); tID != nil {
		resp["term_id"] = *tID
	}
	response.Success(w, resp, "")
}

// POST /auth/logout
// FIX: blacklists the refresh token so it cannot be reused even if
// it has not expired. Access token (15m TTL) is discarded client-side.
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dtos.RefreshRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	if strings.TrimSpace(req.RefreshToken) != "" {
		userID := mw.GetUserID(r.Context())
		_ = h.svc.Logout(req.RefreshToken, userID)
	}
	response.Success(w, nil, "logged out successfully")
}

func sanitiseDecodeError(err error) string {
	if err == nil { return "" }
	s := err.Error()
	if strings.Contains(s, "unknown field") { return s }
	return "invalid JSON"
}

func formatValidationErrors(errs map[string]string) string {
	keys := make([]string, 0, len(errs))
	for k := range errs { keys = append(keys, k) }
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys { parts = append(parts, fmt.Sprintf("%s: %s", k, errs[k])) }
	return strings.Join(parts, "; ")
}
