package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	mw      "school-ms/internal/middleware"
	dtos    "school-ms/internal/Modules/Students/DTOs"
	services "school-ms/internal/Modules/Students/Services"
	"school-ms/internal/pkg/audit"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type StudentHandler struct {
	svc *services.StudentService
	db  *sqlx.DB
}

func NewStudentHandler(svc *services.StudentService, db *sqlx.DB) *StudentHandler {
	return &StudentHandler{svc: svc, db: db}
}

// POST /students
func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil { response.Forbidden(w, "no school context"); return }

	var dto dtos.CreateStudentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	dto.SchoolID = *schoolID
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs)); return
	}
	st, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }

	audit.Log(r.Context(), h.db,
		mw.GetTenantID(r.Context()), *schoolID, mw.GetUserID(r.Context()),
		"create", "student", st.ID,
		map[string]interface{}{"admission_no": st.AdmissionNo})

	response.Created(w, st, "student enrolled")
}

// GET /students  — paginated for admin; filtered by parent/teacher role
func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	roles    := mw.GetRoles(r.Context()) // []string — GetRoles exists; GetRole does not
	userID   := mw.GetUserID(r.Context())

	hasRole := func(role string) bool {
		for _, ro := range roles { if ro == role { return true } }
		return false
	}

	page, _    := strconv.Atoi(r.URL.Query().Get("page"));     if page < 1    { page = 1 }
	perPage, _ := strconv.Atoi(r.URL.Query().Get("per_page")); if perPage < 1 { perPage = 50 }

	switch {
	case hasRole("parent"):
		list, err := h.svc.ListByParentUser(userID)
		if err != nil { response.InternalError(w, err.Error()); return }
		response.Success(w, list, "")
	case hasRole("teacher"):
		list, err := h.svc.ListByTeacherUser(userID)
		if err != nil { response.InternalError(w, err.Error()); return }
		response.Success(w, list, "")
	default:
		if schoolID == nil { response.Success(w, []interface{}{}, ""); return }
		list, total, err := h.svc.List(*schoolID, page, perPage)
		if err != nil { response.InternalError(w, err.Error()); return }
		totalPages := int(total) / perPage
		if int(total)%perPage != 0 { totalPages++ }
		response.Paginated(w, list, response.Meta{
			Page: page, PerPage: perPage, Total: total, TotalPages: totalPages,
		})
	}
}

// GET /students/search?q=
func (h *StudentHandler) Search(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil { response.Success(w, []interface{}{}, ""); return }
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if len(q) < 2 {
		response.BadRequest(w, "search query must be at least 2 characters"); return
	}
	list, err := h.svc.Search(*schoolID, q)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

// GET /students/{id}
func (h *StudentHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid student id"); return }

	roles  := mw.GetRoles(r.Context())
	userID := mw.GetUserID(r.Context())

	hasRole := func(role string) bool {
		for _, ro := range roles { if ro == role { return true } }
		return false
	}

	st, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "student not found"); return }

	switch {
	case hasRole("parent"):
		if linked, _ := h.svc.IsParentOfStudent(userID, id); !linked {
			response.Forbidden(w, "you do not have access to this student"); return
		}
	case hasRole("teacher"):
		if linked, _ := h.svc.IsTeacherOfStudent(userID, id); !linked {
			response.Forbidden(w, "you do not have access to this student"); return
		}
	}
	response.Success(w, st, "")
}

// GET /students/class/{classId}
func (h *StudentHandler) ListByClass(w http.ResponseWriter, r *http.Request) {
	classID, err := strconv.ParseInt(chi.URLParam(r, "classId"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid class id"); return }
	list, err := h.svc.ListByClass(classID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

// PUT /students/{id}
func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid student id"); return }

	var dto dtos.UpdateStudentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs)); return
	}
	if err := h.svc.Update(id, dto); err != nil {
		response.InternalError(w, err.Error()); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	var sid int64
	if schoolID != nil { sid = *schoolID }
	audit.Log(r.Context(), h.db,
		mw.GetTenantID(r.Context()), sid, mw.GetUserID(r.Context()),
		"update", "student", id, dto)
	response.Success(w, nil, "student updated")
}

// DELETE /students/{id}  — soft-delete
func (h *StudentHandler) Deactivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid student id"); return }
	actorID := mw.GetUserID(r.Context())
	if err := h.svc.Deactivate(id, actorID); err != nil {
		response.InternalError(w, err.Error()); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	var sid int64
	if schoolID != nil { sid = *schoolID }
	audit.Log(r.Context(), h.db,
		mw.GetTenantID(r.Context()), sid, actorID,
		"deactivate", "student", id, nil)
	response.Success(w, nil, "student deactivated")
}

// GET /students/{id}/parents
func (h *StudentHandler) GetParents(w http.ResponseWriter, r *http.Request) {
	studentID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil { response.BadRequest(w, "invalid student id"); return }
	list, err := h.svc.GetParents(studentID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

// POST /students/promote
func (h *StudentHandler) Promote(w http.ResponseWriter, r *http.Request) {
	var dto dtos.PromoteStudentsDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body"); return
	}
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs)); return
	}
	actorID := mw.GetUserID(r.Context())
	count, err := h.svc.Promote(dto, actorID)
	if err != nil { response.InternalError(w, err.Error()); return }

	schoolID := mw.GetSchoolID(r.Context())
	var sid int64
	if schoolID != nil { sid = *schoolID }
	audit.Log(r.Context(), h.db,
		mw.GetTenantID(r.Context()), sid, actorID,
		"promote", "students", int64(dto.ToClassID), dto)
	response.Success(w, map[string]int{"promoted": count}, "students promoted")
}

func formatErrors(errs map[string]string) string {
	keys := make([]string, 0, len(errs))
	for k := range errs { keys = append(keys, k) }
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys { parts = append(parts, fmt.Sprintf("%s: %s", k, errs[k])) }
	return strings.Join(parts, "; ")
}