package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Assignments/DTOs"
	svc  "school-ms/internal/Modules/Assignments/Services"
	"school-ms/internal/pkg/response"
	"github.com/go-chi/chi/v5"
)

type AssignmentHandler struct{ svc *svc.AssignmentService }
func NewAssignmentHandler(s *svc.AssignmentService) *AssignmentHandler { return &AssignmentHandler{svc:s} }

func (h *AssignmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateAssignmentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w,"invalid payload"); return }
	// If teacher is not provided, use the authenticated teacher
	if dto.TeacherID == 0 { dto.TeacherID = mw.GetUserID(r.Context()) }
	a,err := h.svc.Create(dto,mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Created(w,a,"assignment created")
}
func (h *AssignmentHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	termID,_ := strconv.ParseInt(r.URL.Query().Get("term_id"),10,64)
	classID,_ := strconv.ParseInt(r.URL.Query().Get("class_id"),10,64)
	if classID > 0 {
		list,err := h.svc.ListByClass(classID,termID)
		if err != nil { response.InternalError(w,err.Error()); return }
		response.Success(w,list,""); return
	}
	list,err := h.svc.ListBySchool(schoolID,termID)
	if err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,list,"")
}
func (h *AssignmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id,_ := strconv.ParseInt(chi.URLParam(r,"id"),10,64)
	if err := h.svc.Delete(id); err != nil { response.InternalError(w,err.Error()); return }
	response.Success(w,nil,"assignment deleted")
}
