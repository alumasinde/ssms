package handlers

import (
	"encoding/json"
	"errors"
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
//
// Public. ResolveTenantFromDB middleware must run before this handler.
//
// Request:  {"email":"...","password":"..."}
// Response: LoginResponse — token pair, user summary, roles, permissions, context.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 1. Tenant guard
	tenant := mw.GetResolvedTenant(r.Context())
	if tenant.ID == 0 {
		response.Unauthorized(w,
			"tenant could not be resolved — check Host header or X-Forwarded-Host")
		return
	}

	// 2. Parse — DisallowUnknownFields rejects typos like "pasword"
	defer r.Body.Close()
	var req dtos.LoginRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		response.BadRequest(w, "invalid JSON: "+sanitiseDecodeError(err))
		return
	}

	// 3. Normalize (lowercase + trim) then field-level validation
	req.Normalize()
	if errs := req.Validate(); errs != nil {
		response.BadRequest(w, formatValidationErrors(errs))
		return
	}

	// 4. Delegate — all HTTP translation lives here, not in the service
	res, err := h.svc.Login(req, tenant.ID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrAccountDisabled):
			// 403: credentials correct but account blocked
			response.Forbidden(w, "this account has been disabled — contact your administrator")
		default:
			// ErrInvalidCredentials and any unexpected internal error → 401
			// Never hint whether the email exists.
			response.Unauthorized(w, "invalid email or password")
		}
		return
	}

	response.Success(w, res, "login successful")
}

// POST /auth/refresh
//
// Public. Validates a refresh token and returns a new access + refresh pair.
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
		response.Unauthorized(w, err.Error())
		return
	}
	response.Success(w, res, "token refreshed")
}

// GET /auth/me
//
// Protected — Authenticate middleware must run first.
// Returns caller's identity from token claims (zero DB queries).
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp := map[string]interface{}{
		"user_id":   mw.GetUserID(ctx),
		"tenant_id": mw.GetTenantID(ctx),
		"school_id": mw.GetSchoolID(ctx), // *int64, nil for superadmin
		"roles":     mw.GetRoles(ctx),    // []string
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
//
// Protected. For stateless JWT, logout is client-side (discard the token).
// Endpoint is a hook for future server-side revocation via Redis JTI blocklist.
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: push token JTI to Redis with TTL = ExpiresAt - now to block reuse.
	response.Success(w, nil, "logged out successfully")
}

// ── Private helpers ───────────────────────────────────────────────────────────

func sanitiseDecodeError(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	if strings.Contains(s, "unknown field") {
		return s
	}
	return "invalid JSON"
}

func formatValidationErrors(errs map[string]string) string {
	keys := make([]string, 0, len(errs))
	for k := range errs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s: %s", k, errs[k]))
	}
	return strings.Join(parts, "; ")
}
