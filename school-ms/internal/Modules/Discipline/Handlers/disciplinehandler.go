package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Discipline/DTOs"
	svc  "school-ms/internal/Modules/Discipline/Services"
	"school-ms/internal/pkg/response"
	"github.com/go-chi/chi/v5"
)

type DisciplineHandler struct{ svc *svc.DisciplineService }
func NewDisciplineHandler(s *svc.DisciplineService) *DisciplineHandler { return &DisciplineHandler{svc:s} }

func (h *DisciplineHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateDisciplineDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w,"invalid payload"); return }
	d,err := h.svc.Create(dto,mw.GetSchoolID(r.Context()),mw.GetUserID(r.Context()))
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Created(w,d,"record created")
}
func (h *DisciplineHandler) List(w http.ResponseWriter, r *http.Request) {
	termID,_ := strconv.ParseInt(r.URL.Query().Get("term_id"),10,64)
	list,err := h.svc.ListBySchool(mw.GetSchoolID(r.Context()),termID)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
func (h *DisciplineHandler) ListByStudent(w http.ResponseWriter, r *http.Request) {
	studentID,_ := strconv.ParseInt(chi.URLParam(r,"studentId"),10,64)
	list,err := h.svc.ListByStudent(studentID)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
func (h *DisciplineHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id,_ := strconv.ParseInt(chi.URLParam(r,"id"),10,64)
	if err := h.svc.Delete(id); err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,nil,"record deleted")
}
