package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Parents/DTOs"
	services "school-ms/internal/Modules/Parents/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ParentHandler struct{ svc *services.ParentService }

func NewParentHandler(svc *services.ParentService) *ParentHandler { return &ParentHandler{svc: svc} }

func (h *ParentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateParentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	p, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, p, "parent created")
}

func (h *ParentHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.List(mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *ParentHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	p, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "parent not found"); return }
	response.Success(w, p, "")
}

func (h *ParentHandler) LinkStudent(w http.ResponseWriter, r *http.Request) {
	var dto dtos.LinkStudentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	if err := h.svc.LinkStudent(dto); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "student linked to parent")
}

func (h *ParentHandler) GetStudentParents(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	list, err := h.svc.GetStudentParents(studentID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}
