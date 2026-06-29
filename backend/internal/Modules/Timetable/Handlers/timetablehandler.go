package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Timetable/DTOs"
	svc  "school-ms/internal/Modules/Timetable/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type TimetableHandler struct{ svc *svc.TimetableService }

func NewTimetableHandler(s *svc.TimetableService) *TimetableHandler { return &TimetableHandler{svc: s} }

// POST /timetable — service Create(dto, schoolID int64), nil guard + dereference
func (h *TimetableHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateSlotDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	slot, err := h.svc.Create(dto, *schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, slot, "slot created")
}

// PUT /timetable/{id} — no schoolID from context, correct as-is
func (h *TimetableHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.CreateSlotDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	if err := h.svc.Update(id, dto); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "slot updated")
}

// DELETE /timetable/{id} — no schoolID from context, correct as-is
func (h *TimetableHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "slot deleted")
}

// GET /timetable/class/{classId} — URL params are int64, correct as-is
func (h *TimetableHandler) ListByClass(w http.ResponseWriter, r *http.Request) {
	classID, _ := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	termID, _  := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	list, err  := h.svc.ListByClass(classID, termID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /timetable/teacher/{teacherId} — URL params are int64, correct as-is
func (h *TimetableHandler) ListByTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID, _ := strconv.ParseInt(chi.URLParam(r, "teacherId"), 10, 64)
	termID, _    := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	list, err    := h.svc.ListByTeacher(teacherID, termID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /timetable/school — service ListBySchool(schoolID int64, termID int64)
// nil guard + dereference required
func (h *TimetableHandler) ListBySchool(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	list, err  := h.svc.ListBySchool(*schoolID, termID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}