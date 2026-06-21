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
func NewStaffAttendanceHandler(s *svc.StaffAttendanceService) *StaffAttendanceHandler { return &StaffAttendanceHandler{svc:s} }

func (h *StaffAttendanceHandler) Mark(w http.ResponseWriter, r *http.Request) {
	var dto dtos.MarkStaffDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w,"invalid payload"); return }
	if err := h.svc.Mark(dto,mw.GetSchoolID(r.Context()),mw.GetUserID(r.Context())); err != nil {
		response.InternalError(w,err.Error()); return
	}
	response.Success(w,nil,"staff attendance marked")
}
func (h *StaffAttendanceHandler) ListByDate(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	list,err := h.svc.ListByDate(mw.GetSchoolID(r.Context()),date)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
func (h *StaffAttendanceHandler) Summary(w http.ResponseWriter, r *http.Request) {
	month := r.URL.Query().Get("month")
	list,err := h.svc.Summary(mw.GetSchoolID(r.Context()),month)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
