package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Users/DTOs"
	services "school-ms/internal/Modules/Users/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var dto dtos.RegisterDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload")
		return
	}
	u, err := h.svc.Register(dto)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Created(w, dtos.UserResponse{
		ID: u.ID, Name: u.Name, Email: u.Email, Role: u.Role, TenantID: u.TenantID,
	}, "user registered")
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	u, err := h.svc.GetByID(mw.GetUserID(r.Context()))
	if err != nil {
		response.NotFound(w, "user not found")
		return
	}
	response.Success(w, dtos.UserResponse{
		ID: u.ID, Name: u.Name, Email: u.Email, Role: u.Role, TenantID: u.TenantID,
	}, "")
}

func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var dto dtos.ChangePasswordDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload")
		return
	}
	if err := h.svc.ChangePassword(mw.GetUserID(r.Context()), dto); err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Success(w, nil, "password changed")
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.ListByTenant(mw.GetTenantID(r.Context()))
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, list, "")
}

func (h *UserHandler) Deactivate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Deactivate(id); err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, nil, "user deactivated")
}
