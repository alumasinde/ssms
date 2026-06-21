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
func NewTimetableHandler(s *svc.TimetableService) *TimetableHandler { return &TimetableHandler{svc:s} }

func (h *TimetableHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateSlotDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w,"invalid payload"); return }
	slot,err := h.svc.Create(dto,mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Created(w,slot,"slot created")
}
func (h *TimetableHandler) Update(w http.ResponseWriter, r *http.Request) {
	id,_ := strconv.ParseInt(chi.URLParam(r,"id"),10,64)
	var dto dtos.CreateSlotDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w,"invalid payload"); return }
	if err := h.svc.Update(id,dto); err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,nil,"slot updated")
}
func (h *TimetableHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id,_ := strconv.ParseInt(chi.URLParam(r,"id"),10,64)
	if err := h.svc.Delete(id); err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,nil,"slot deleted")
}
func (h *TimetableHandler) ListByClass(w http.ResponseWriter, r *http.Request) {
	classID,_ := strconv.ParseInt(chi.URLParam(r,"classId"),10,64)
	termID,_  := strconv.ParseInt(r.URL.Query().Get("term_id"),10,64)
	list,err  := h.svc.ListByClass(classID,termID)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
func (h *TimetableHandler) ListByTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID,_ := strconv.ParseInt(chi.URLParam(r,"teacherId"),10,64)
	termID,_    := strconv.ParseInt(r.URL.Query().Get("term_id"),10,64)
	list,err    := h.svc.ListByTeacher(teacherID,termID)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
func (h *TimetableHandler) ListBySchool(w http.ResponseWriter, r *http.Request) {
	termID,_ := strconv.ParseInt(r.URL.Query().Get("term_id"),10,64)
	list,err := h.svc.ListBySchool(mw.GetSchoolID(r.Context()),termID)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
