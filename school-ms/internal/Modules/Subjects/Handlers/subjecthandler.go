package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Subjects/DTOs"
	services "school-ms/internal/Modules/Subjects/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type SubjectHandler struct{ svc *services.SubjectService }

func NewSubjectHandler(svc *services.SubjectService) *SubjectHandler { return &SubjectHandler{svc: svc} }

func (h *SubjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateSubjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	s, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, s, "subject created")
}

func (h *SubjectHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.List(mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *SubjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	s, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "subject not found"); return }
	response.Success(w, s, "")
}

func (h *SubjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.CreateSubjectDTO
	json.NewDecoder(r.Body).Decode(&dto)
	if err := h.svc.Update(id, dto); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "subject updated")
}

func (h *SubjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "subject deleted")
}
