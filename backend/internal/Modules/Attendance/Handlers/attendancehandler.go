package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Attendance/DTOs"
	services "school-ms/internal/Modules/Attendance/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type AttendanceHandler struct{ svc *services.AttendanceService }

func NewAttendanceHandler(svc *services.AttendanceService) *AttendanceHandler { return &AttendanceHandler{svc: svc} }

func (h *AttendanceHandler) Mark(w http.ResponseWriter, r *http.Request) {
	var dto dtos.MarkAttendanceDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	if err := h.svc.Mark(dto, mw.GetUserID(r.Context())); err != nil { response.ServerError(w, err); return }
	response.Success(w, nil, "attendance recorded")
}

func (h *AttendanceHandler) GetByClassDate(w http.ResponseWriter, r *http.Request) {
	classID, _ := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	date := r.URL.Query().Get("date")
	list, err := h.svc.GetByClassDate(classID, date)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

func (h *AttendanceHandler) GetByStudent(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	list, err := h.svc.GetByStudent(studentID, termID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

func (h *AttendanceHandler) Summary(w http.ResponseWriter, r *http.Request) {
	classID, _ := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	list, err := h.svc.GetSummary(classID, termID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}
