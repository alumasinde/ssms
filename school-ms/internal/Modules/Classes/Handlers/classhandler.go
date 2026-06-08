package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Classes/DTOs"
	services "school-ms/internal/Modules/Classes/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ClassHandler struct{ svc *services.ClassService }

func NewClassHandler(svc *services.ClassService) *ClassHandler { return &ClassHandler{svc: svc} }

func (h *ClassHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateClassDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	c, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, c, "class created")
}

func (h *ClassHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	list, err := h.svc.List(schoolID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *ClassHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	c, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "class not found"); return }
	response.Success(w, c, "")
}

func (h *ClassHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.CreateClassDTO
	json.NewDecoder(r.Body).Decode(&dto)
	if err := h.svc.Update(id, dto); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "class updated")
}

func (h *ClassHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "class deleted")
}
