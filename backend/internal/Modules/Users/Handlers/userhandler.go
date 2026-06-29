package handlers
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	mw     "school-ms/internal/middleware"
	dtos   "school-ms/internal/Modules/Users/DTOs"
	models "school-ms/internal/Modules/Users/Models"
	svc    "school-ms/internal/Modules/Users/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct{ svc *svc.UserService }

func NewUserHandler(s *svc.UserService) *UserHandler { return &UserHandler{svc: s} }

// POST /users
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	tenantID := mw.GetTenantID(r.Context())
	schoolID := mw.GetSchoolID(r.Context()) // *int64

	var dto dtos.RegisterDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	dto.Normalize()
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs)); return
	}
	// Inject school context for non-superadmin callers
	if dto.SchoolID == nil && schoolID != nil {
		dto.SchoolID = schoolID
	}
	u, err := h.svc.Register(dto, tenantID, dto.SchoolID)
	if err != nil {
		switch {
		case errors.Is(err, svc.ErrEmailTaken):
			response.BadRequest(w, err.Error())
		case errors.Is(err, svc.ErrInvalidRole):
			response.BadRequest(w, "role_code is invalid or inactive")
		default:
			response.ServerError(w, err)
		}
		return
	}
	response.Created(w, toResponse(u), "user registered")
}

// GET /users/me
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	u, err := h.svc.GetByID(mw.GetUserID(r.Context()))
	if err != nil { response.NotFound(w, "user not found"); return }
	response.Success(w, toResponse(u), "")
}

// GET /users/{id}
// FIX: enforce tenant isolation — a user from tenant A must not be able
// to fetch a user from tenant B just by guessing an integer ID.
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid user id"); return }
	u, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "user not found"); return }
	// Tenant isolation check
	if u.TenantID != mw.GetTenantID(r.Context()) {
		response.NotFound(w, "user not found"); return
	}
	response.Success(w, toResponse(u), "")
}

// GET /users  (tenant-wide, superadmin)
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.ListByTenant(mw.GetTenantID(r.Context()))
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, toResponseList(list), "")
}

// GET /users/school  (school-scoped, admin)
func (h *UserHandler) ListBySchool(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil { response.Success(w, []dtos.UserResponse{}, ""); return }
	list, err := h.svc.ListBySchool(*schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, toResponseList(list), "")
}

// PUT /users/me/password
func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var dto dtos.ChangePasswordDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs)); return
	}
	if err := h.svc.ChangePassword(mw.GetUserID(r.Context()), dto); err != nil {
		if errors.Is(err, svc.ErrWrongPassword) {
			response.BadRequest(w, "current password is incorrect"); return
		}
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "password changed")
}

// PUT /users/{id}/role
func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid user id"); return }
	var dto dtos.UpdateRoleDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	if strings.TrimSpace(dto.RoleCode) == "" {
		response.BadRequest(w, "role_code is required"); return
	}
	if err := h.svc.UpdateRole(id, mw.GetTenantID(r.Context()), dto); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "role updated")
}

// POST /users/{id}/activate
func (h *UserHandler) Activate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid user id"); return }
	if err := h.svc.Activate(id); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "user activated")
}

// DELETE /users/{id}  — soft-delete
func (h *UserHandler) Deactivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid user id"); return }
	actorID := mw.GetUserID(r.Context())
	if id == actorID {
		response.BadRequest(w, "you cannot deactivate your own account"); return
	}
	if err := h.svc.Deactivate(id, actorID); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "user deactivated")
}

// ── helpers ────────────────────────────────────────────────────────────

func toResponse(u *models.UserWithRoles) dtos.UserResponse {
	return dtos.UserResponse{
		ID: u.ID, FirstName: u.FirstName, LastName: u.LastName,
		Name: u.FullName(), Email: u.Email, Phone: u.Phone,
		SchoolID: u.SchoolID, TenantID: u.TenantID,
		Roles: u.Roles, IsActive: u.IsActive,
	}
}

func toResponseList(list []models.UserWithRoles) []dtos.UserResponse {
	out := make([]dtos.UserResponse, 0, len(list))
	for i := range list { out = append(out, toResponse(&list[i])) }
	return out
}

func formatErrors(errs map[string]string) string {
	keys := make([]string, 0, len(errs))
	for k := range errs { keys = append(keys, k) }
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys { parts = append(parts, fmt.Sprintf("%s: %s", k, errs[k])) }
	return strings.Join(parts, "; ")
}