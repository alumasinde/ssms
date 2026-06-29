package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw       "school-ms/internal/middleware"
	dtos     "school-ms/internal/Modules/Subjects/DTOs"
	services "school-ms/internal/Modules/Subjects/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type SubjectHandler struct{ svc *services.SubjectService }

func NewSubjectHandler(svc *services.SubjectService) *SubjectHandler { return &SubjectHandler{svc: svc} }

// POST /subjects
// dto.SchoolID is int64 — nil guard + dereference required
func (h *SubjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateSubjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	dto.SchoolID = *schoolID
	s, err := h.svc.Create(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, s, "subject created")
}

// GET /subjects
// service List(schoolID int64) — nil guard + dereference required
func (h *SubjectHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	list, err := h.svc.List(*schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /subjects/{id}
// no schoolID involved — correct as-is
func (h *SubjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	s, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "subject not found"); return }
	response.Success(w, s, "")
}

// PUT /subjects/{id}
// no schoolID involved — fixed missing error check on Decode
func (h *SubjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.CreateSubjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	if err := h.svc.Update(id, dto); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "subject updated")
}

// DELETE /subjects/{id}
// no schoolID involved — correct as-is
func (h *SubjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "subject deleted")
}
