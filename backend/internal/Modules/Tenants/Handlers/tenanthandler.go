package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	dtos "school-ms/internal/Modules/Tenants/DTOs"
	services "school-ms/internal/Modules/Tenants/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type TenantHandler struct {
	svc *services.TenantService
}

func NewTenantHandler(svc *services.TenantService) *TenantHandler {
	return &TenantHandler{svc: svc}
}

func (h *TenantHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateTenantDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload")
		return
	}
	t, err := h.svc.Create(dto)
	if err != nil {
		response.ServerError(w, err)
		return
	}
	response.Created(w, t, "tenant created")
}

func (h *TenantHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.List()
	if err != nil {
		response.ServerError(w, err)
		return
	}
	response.Success(w, list, "")
}

func (h *TenantHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	t, err := h.svc.GetByID(id)
	if err != nil {
		response.NotFound(w, "tenant not found")
		return
	}
	response.Success(w, t, "")
}

func (h *TenantHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.UpdateTenantDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload")
		return
	}
	if err := h.svc.Update(id, dto); err != nil {
		response.ServerError(w, err)
		return
	}
	response.Success(w, nil, "tenant updated")
}

func (h *TenantHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(w, err)
		return
	}
	response.Success(w, nil, "tenant deactivated")
}
