package handlers

import (
	"encoding/json"
	"net/http"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/StaffAttendance/DTOs"
	svc  "school-ms/internal/Modules/StaffAttendance/Services"
	"school-ms/internal/pkg/response"
)

type StaffAttendanceHandler struct{ svc *svc.StaffAttendanceService }

func NewStaffAttendanceHandler(s *svc.StaffAttendanceService) *StaffAttendanceHandler {
	return &StaffAttendanceHandler{svc: s}
}

// POST /staff-attendance
// service Mark(dto, schoolID int64, recordedBy int64)
// nil guard + dereference required + InternalError → ServerError
func (h *StaffAttendanceHandler) Mark(w http.ResponseWriter, r *http.Request) {
	var dto dtos.MarkStaffDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	if err := h.svc.Mark(dto, *schoolID, mw.GetUserID(r.Context())); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "staff attendance marked")
}

// GET /staff-attendance?date=
// service ListByDate(schoolID int64, date string)
// nil guard + dereference required + InternalError → ServerError
func (h *StaffAttendanceHandler) ListByDate(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	date := r.URL.Query().Get("date")
	list, err := h.svc.ListByDate(*schoolID, date)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /staff-attendance/summary?month=
// service Summary(schoolID int64, yearMonth string)
// nil guard + dereference required + InternalError → ServerError
func (h *StaffAttendanceHandler) Summary(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	month := r.URL.Query().Get("month")
	list, err := h.svc.Summary(*schoolID, month)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}
