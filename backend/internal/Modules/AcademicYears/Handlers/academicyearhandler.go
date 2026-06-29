package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/AcademicYears/DTOs"
	services "school-ms/internal/Modules/AcademicYears/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type AcademicYearHandler struct{ svc *services.AcademicYearService }

func NewAcademicYearHandler(svc *services.AcademicYearService) *AcademicYearHandler {
	return &AcademicYearHandler{svc: svc}
}

func (h *AcademicYearHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateAcademicYearDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	dto.SchoolID = *schoolID
	ay, err := h.svc.Create(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, ay, "academic year created")
}

func (h *AcademicYearHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	list, err := h.svc.List(*schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

func (h *AcademicYearHandler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.NotFound(w, "no current academic year set"); return
	}
	ay, err := h.svc.GetCurrent(*schoolID)
	if err != nil { response.NotFound(w, "no current academic year set"); return }
	response.Success(w, ay, "")
}

func (h *AcademicYearHandler) SetCurrent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	if err := h.svc.SetCurrent(*schoolID, id); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "current academic year updated")
}

// Get, Update, Delete are fine — no schoolID involved

func (h *AcademicYearHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	ay, err := h.svc.GetByID(id)
	if err != nil {
		response.NotFound(w, "academic year not found")
		return
	}
	response.Success(w, ay, "")
}

func (h *AcademicYearHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.CreateAcademicYearDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload")
		return
	}
	if err := h.svc.Update(id, dto); err != nil {
		response.ServerError(w, err)
		return
	}
	response.Success(w, nil, "academic year updated")
}

func (h *AcademicYearHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(w, err)
		return
	}
	response.Success(w, nil, "academic year deleted")
}
