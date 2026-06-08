package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Students/DTOs"
	services "school-ms/internal/Modules/Students/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type StudentHandler struct{ svc *services.StudentService }

func NewStudentHandler(svc *services.StudentService) *StudentHandler { return &StudentHandler{svc: svc} }

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateStudentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	st, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, st, "student enrolled")
}

func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	role     := mw.GetRole(r.Context())
	userID   := mw.GetUserID(r.Context())

	page, _    := strconv.Atoi(r.URL.Query().Get("page"));     if page < 1    { page = 1 }
	perPage, _ := strconv.Atoi(r.URL.Query().Get("per_page")); if perPage < 1 { perPage = 50 }

	switch role {
	case "parent":
		// Parent sees only their linked children
		list, err := h.svc.ListByParentUser(userID)
		if err != nil { response.InternalError(w, err.Error()); return }
		response.Success(w, list, "")

	case "teacher":
		// Teacher sees only students in their assigned classes
		list, err := h.svc.ListByTeacherUser(userID)
		if err != nil { response.InternalError(w, err.Error()); return }
		response.Success(w, list, "")

	default:
		// Admin, superadmin see all students in school
		list, total, err := h.svc.List(schoolID, page, perPage)
		if err != nil { response.InternalError(w, err.Error()); return }
		totalPages := int(total) / perPage
		if int(total)%perPage != 0 { totalPages++ }
		response.Paginated(w, list, response.Meta{Page: page, PerPage: perPage, Total: total, TotalPages: totalPages})
	}
}

func (h *StudentHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _  := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	role   := mw.GetRole(r.Context())
	userID := mw.GetUserID(r.Context())

	st, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "student not found"); return }

	switch role {
	case "parent":
		linked, err := h.svc.IsParentOfStudent(userID, id)
		if err != nil || !linked {
			response.Forbidden(w, "you do not have access to this student")
			return
		}
	case "teacher":
		// Teacher can only view students in their classes
		linked, err := h.svc.IsTeacherOfStudent(userID, id)
		if err != nil || !linked {
			response.Forbidden(w, "you do not have access to this student")
			return
		}
	}

	response.Success(w, st, "")
}

func (h *StudentHandler) ListByClass(w http.ResponseWriter, r *http.Request) {
	classID, _ := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	list, err := h.svc.ListByClass(classID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.CreateStudentDTO
	json.NewDecoder(r.Body).Decode(&dto)
	if err := h.svc.Update(id, dto); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "student updated")
}

func (h *StudentHandler) Deactivate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Deactivate(id); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "student deactivated")
}
