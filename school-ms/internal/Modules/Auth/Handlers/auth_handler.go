package handlers

import (
	"encoding/json"
	"net/http"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Auth/DTOs"
	"school-ms/internal/Modules/Auth/Services"
	"school-ms/internal/pkg/response"
)

type AuthHandler struct {
	svc *services.AuthService
}

func NewAuthHandler(svc *services.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// Login accepts { "email": "...", "password": "..." } — NO tenant_id.
// The tenant is resolved from the HTTP Host header by the ResolveTenantFromDB middleware.
//
// Supported URL patterns:
//   ssms.highwayhighschool.ac.ke  →  subdomain slug = "highwayhighschool"
//   highwayhighschool.ac.ke       →  direct school domain match
//   highwayhighschool.ac.ke/ssms  →  path prefix (handled by router mount)
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Tenant was resolved by ResolveTenantFromDB middleware
	tenant := mw.GetResolvedTenant(r.Context())

	var req dtos.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		response.BadRequest(w, "email and password are required")
		return
	}

	res, err := h.svc.Login(req, tenant.ID)
	if err != nil {
		response.Unauthorized(w, err.Error())
		return
	}

	response.Success(w, res, "login successful")
}

// Whoami returns the current authenticated user's claims – useful for the
// frontend to re-hydrate session state without a round-trip to /users/me.
func (h *AuthHandler) Whoami(w http.ResponseWriter, r *http.Request) {
	response.Success(w, map[string]interface{}{
		"user_id":   mw.GetUserID(r.Context()),
		"tenant_id": mw.GetTenantID(r.Context()),
		"school_id": mw.GetSchoolID(r.Context()),
		"role":      mw.GetRole(r.Context()),
	}, "")
}
