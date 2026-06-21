package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Classes/DTOs"
	svc  "school-ms/internal/Modules/Classes/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ClassSubjectHandler struct{ svc *svc.ClassSubjectService }

func NewClassSubjectHandler(s *svc.ClassSubjectService) *ClassSubjectHandler {
	return &ClassSubjectHandler{svc: s}
}

// GET /classes/{id}/subjects
func (h *ClassSubjectHandler) List(w http.ResponseWriter, r *http.Request) {
	classID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid class id")
		return
	}

	list, err := h.svc.List(classID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, list, "")
}

// GET /classes/{id}/subjects/unassigned
func (h *ClassSubjectHandler) ListUnassigned(w http.ResponseWriter, r *http.Request) {
	classID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid class id")
		return
	}

	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, "")
		return
	}

	list, err := h.svc.ListUnassigned(classID, *schoolID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, list, "")
}

// POST /classes/{id}/subjects
func (h *ClassSubjectHandler) Assign(w http.ResponseWriter, r *http.Request) {
	classID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid class id")
		return
	}

	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context")
		return
	}

	var dto dtos.AssignSubjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body")
		return
	}
	if dto.SubjectID == 0 {
		response.BadRequest(w, "subject_id is required")
		return
	}

	if err := h.svc.Assign(classID, dto.SubjectID, *schoolID, dto.IsCompulsory); err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, nil, "subject assigned")
}

// DELETE /classes/{id}/subjects/{subjectId}
func (h *ClassSubjectHandler) Remove(w http.ResponseWriter, r *http.Request) {
	classID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid class id")
		return
	}

	subjectID, err := strconv.ParseInt(chi.URLParam(r, "subjectId"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid subject id")
		return
	}

	if err := h.svc.Remove(classID, subjectID); err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, nil, "subject removed")
}
