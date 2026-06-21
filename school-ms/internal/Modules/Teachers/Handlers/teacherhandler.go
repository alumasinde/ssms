package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Teachers/DTOs"
	svc  "school-ms/internal/Modules/Teachers/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type TeacherHandler struct{ svc *svc.TeacherService }

func NewTeacherHandler(s *svc.TeacherService) *TeacherHandler { return &TeacherHandler{svc: s} }

// POST /teachers
func (h *TeacherHandler) Create(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil { response.Forbidden(w, "no school context"); return }
	var dto dtos.CreateTeacherDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	dto.SchoolID = *schoolID
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs)); return
	}
	t, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, t, "teacher created")
}

// GET /teachers
func (h *TeacherHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil { response.Success(w, []interface{}{}, ""); return }
	list, err := h.svc.List(*schoolID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

// GET /teachers/{id}
func (h *TeacherHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid teacher id"); return }
	t, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "teacher not found"); return }
	response.Success(w, t, "")
}

// PUT /teachers/{id}  — patch semantics, ID from URL param only
func (h *TeacherHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid teacher id"); return }
	var dto dtos.UpdateTeacherDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	updated, err := h.svc.Update(id, dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, updated, "teacher updated")
}

// DELETE /teachers/{id}  — soft-delete
func (h *TeacherHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid teacher id"); return }
	if err := h.svc.Delete(id, mw.GetUserID(r.Context())); err != nil {
		response.InternalError(w, err.Error()); return
	}
	response.Success(w, nil, "teacher deleted")
}

// POST /teachers/{id}/subjects  — teacher_id from URL, NOT body
func (h *TeacherHandler) AssignSubject(w http.ResponseWriter, r *http.Request) {
	teacherID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid teacher id"); return }
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil { response.Forbidden(w, "no school context"); return }
	var dto dtos.AssignSubjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs)); return
	}
	if err := h.svc.AssignSubject(teacherID, *schoolID, dto); err != nil {
		response.InternalError(w, err.Error()); return
	}
	response.Success(w, nil, "subject assigned")
}

// GET /teachers/{id}/subjects
func (h *TeacherHandler) GetSubjects(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid teacher id"); return }
	list, err := h.svc.GetSubjects(id)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

// DELETE /teachers/{id}/subjects/{subjectId}/{classId}
func (h *TeacherHandler) RemoveSubject(w http.ResponseWriter, r *http.Request) {
	teacherID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid teacher id"); return }
	subjectID, err := strconv.ParseInt(chi.URLParam(r, "subjectId"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid subject id"); return }
	classID, err := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid class id"); return }
	if err := h.svc.RemoveSubject(teacherID, subjectID, classID); err != nil {
		response.InternalError(w, err.Error()); return
	}
	response.Success(w, nil, "subject removed")
}

func formatErrors(errs map[string]string) string {
	keys := make([]string, 0, len(errs))
	for k := range errs { keys = append(keys, k) }
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys { parts = append(parts, fmt.Sprintf("%s: %s", k, errs[k])) }
	return strings.Join(parts, "; ")
}