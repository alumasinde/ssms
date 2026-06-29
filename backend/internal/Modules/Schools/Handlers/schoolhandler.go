package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw       "school-ms/internal/middleware"
	dtos     "school-ms/internal/Modules/Schools/DTOs"
	services "school-ms/internal/Modules/Schools/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type SchoolHandler struct {
	svc *services.SchoolService
}

func NewSchoolHandler(svc *services.SchoolService) *SchoolHandler {
	return &SchoolHandler{svc: svc}
}

// POST /schools
// no schoolID from context — tenantID is int64 from GetTenantID, correct as-is
func (h *SchoolHandler) CreateSchool(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateSchoolDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	school, err := h.svc.CreateSchool(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, school, "school created")
}

// GET /schools
// GetTenantID returns int64 directly — correct as-is
func (h *SchoolHandler) ListSchools(w http.ResponseWriter, r *http.Request) {
	tenantID := mw.GetTenantID(r.Context())
	list, err := h.svc.ListSchools(tenantID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /schools/{id}
// no schoolID from context — correct as-is
func (h *SchoolHandler) GetSchool(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	school, err := h.svc.GetSchool(id)
	if err != nil { response.NotFound(w, "school not found"); return }
	response.Success(w, school, "")
}

// POST /schools/{id}/academic-years
// no schoolID from context — schoolID comes from DTO body, correct as-is
func (h *SchoolHandler) CreateAcademicYear(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateAcademicYearDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	ay, err := h.svc.CreateAcademicYear(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, ay, "academic year created")
}

// GET /schools/{schoolId}/academic-years
// schoolID from URL param as int64 — correct as-is
func (h *SchoolHandler) ListAcademicYears(w http.ResponseWriter, r *http.Request) {
	schoolID, _ := strconv.ParseInt(chi.URLParam(r, "schoolId"), 10, 64)
	list, err := h.svc.ListAcademicYears(schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// POST /schools/{id}/terms
// no schoolID from context — schoolID comes from DTO body, correct as-is
func (h *SchoolHandler) CreateTerm(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateTermDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	t, err := h.svc.CreateTerm(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, t, "term created")
}

// GET /schools/{id}/terms/{yearId}
// yearID from URL param as int64 — correct as-is
func (h *SchoolHandler) ListTerms(w http.ResponseWriter, r *http.Request) {
	yearID, _ := strconv.ParseInt(chi.URLParam(r, "yearId"), 10, 64)
	list, err := h.svc.ListTerms(yearID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /schools/current-term
// service GetCurrentTerm(schoolID int64) — nil guard + dereference required
func (h *SchoolHandler) GetCurrentTerm(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.NotFound(w, "no active term found"); return
	}
	t, err := h.svc.GetCurrentTerm(*schoolID)
	if err != nil { response.NotFound(w, "no active term found"); return }
	response.Success(w, t, "")
}

// PUT /schools/{id}
// no schoolID from context — id from URL param, correct as-is
func (h *SchoolHandler) UpdateSchool(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.CreateSchoolDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	if err := h.svc.UpdateSchool(id, dto); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "school updated")
}
