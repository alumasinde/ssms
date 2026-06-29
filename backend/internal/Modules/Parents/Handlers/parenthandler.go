package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw       "school-ms/internal/middleware"
	dtos     "school-ms/internal/Modules/Parents/DTOs"
	services "school-ms/internal/Modules/Parents/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ParentHandler struct{ svc *services.ParentService }

func NewParentHandler(svc *services.ParentService) *ParentHandler { return &ParentHandler{svc: svc} }

// POST /parents
// dto.SchoolID is int64 — nil guard + dereference required
func (h *ParentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateParentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	dto.SchoolID = *schoolID
	p, err := h.svc.Create(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, p, "parent created")
}

// GET /parents
// service List(schoolID int64) — nil guard + dereference required
func (h *ParentHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	list, err := h.svc.List(*schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /parents/{id}
// no schoolID involved — correct as-is
func (h *ParentHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	p, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "parent not found"); return }
	response.Success(w, p, "")
}

// POST /parents/link-student
// no schoolID involved — correct as-is
func (h *ParentHandler) LinkStudent(w http.ResponseWriter, r *http.Request) {
	var dto dtos.LinkStudentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	if err := h.svc.LinkStudent(dto); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "student linked to parent")
}

// GET /parents/student/{studentId}
// no schoolID involved — correct as-is
func (h *ParentHandler) GetStudentParents(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	list, err := h.svc.GetStudentParents(studentID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}
