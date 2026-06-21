package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	mw    "school-ms/internal/middleware"
	dtos  "school-ms/internal/Modules/Classes/DTOs"
	svc   "school-ms/internal/Modules/Classes/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ClassHandler struct{ svc *svc.ClassService }

func NewClassHandler(s *svc.ClassService) *ClassHandler { return &ClassHandler{svc: s} }

// POST /classes
func (h *ClassHandler) Create(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context — cannot create class")
		return
	}

	var dto dtos.CreateClassDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body")
		return
	}
	dto.SchoolID = *schoolID

	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs))
		return
	}

	c, err := h.svc.Create(dto)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Created(w, c, "class created")
}

// GET /classes
func (h *ClassHandler) List(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, "")
		return
	}

	list, err := h.svc.List(*schoolID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, list, "")
}

// GET /classes/{id}
func (h *ClassHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid class id")
		return
	}

	c, err := h.svc.GetByID(id)
	if err != nil {
		response.NotFound(w, "class not found")
		return
	}
	response.Success(w, c, "")
}

// PUT /classes/{id}
func (h *ClassHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid class id")
		return
	}

	var dto dtos.UpdateClassDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid JSON body")
		return
	}
	if errs := dto.Validate(); errs != nil {
		response.BadRequest(w, formatErrors(errs))
		return
	}

	if err := h.svc.Update(id, dto); err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, nil, "class updated")
}

// DELETE /classes/{id}
// Performs a soft-delete — sets deleted_at and deleted_by; preserves FKs.
func (h *ClassHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.BadRequest(w, "invalid class id")
		return
	}

	actorID := mw.GetUserID(r.Context())

	if err := h.svc.Delete(id, actorID); err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, nil, "class deleted")
}

// ── helpers ───────────────────────────────────────────────────────────────────

func formatErrors(errs map[string]string) string {
	keys := make([]string, 0, len(errs))
	for k := range errs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s: %s", k, errs[k]))
	}
	return strings.Join(parts, "; ")
}
