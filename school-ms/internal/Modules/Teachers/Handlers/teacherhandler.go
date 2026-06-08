package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Teachers/DTOs"
	services "school-ms/internal/Modules/Teachers/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type TeacherHandler struct{ svc *services.TeacherService }

func NewTeacherHandler(svc *services.TeacherService) *TeacherHandler { return &TeacherHandler{svc: svc} }

func (h *TeacherHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateTeacherDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	dto.SchoolID = mw.GetSchoolID(r.Context())
	t, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, t, "teacher created")
}

func (h *TeacherHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.List(mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *TeacherHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	t, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "teacher not found"); return }
	response.Success(w, t, "")
}

func (h *TeacherHandler) AssignSubject(w http.ResponseWriter, r *http.Request) {
	var dto dtos.AssignSubjectDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	if err := h.svc.AssignSubject(dto); err != nil {
		response.InternalError(w, err.Error()); return
	}
	response.Success(w, nil, "subject assigned")
}

func (h *TeacherHandler) GetSubjects(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	list, err := h.svc.GetSubjects(id)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}
