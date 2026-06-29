package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw  "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Terms/DTOs"
	svc  "school-ms/internal/Modules/Terms/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type TermHandler struct{ svc *svc.TermService }

func NewTermHandler(s *svc.TermService) *TermHandler { return &TermHandler{svc: s} }

// POST /terms — no schoolID from context, correct as-is
func (h *TermHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateTermDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	t, err := h.svc.Create(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, t, "term created")
}

// GET /terms/year/{yearId} — URL param is int64, correct as-is
func (h *TermHandler) ListByYear(w http.ResponseWriter, r *http.Request) {
	yearID, _ := strconv.ParseInt(chi.URLParam(r, "yearId"), 10, 64)
	list, err := h.svc.ListByYear(yearID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /terms/school — service ListBySchool(int64), nil guard + dereference
func (h *TermHandler) ListBySchool(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	list, err := h.svc.ListBySchool(*schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

// GET /terms/current — service GetCurrent(int64), nil guard + dereference
func (h *TermHandler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.NotFound(w, "no active term found"); return
	}
	t, err := h.svc.GetCurrent(*schoolID)
	if err != nil { response.NotFound(w, "no active term found"); return }
	response.Success(w, t, "")
}

// GET /terms/{id} — no schoolID from context, correct as-is
func (h *TermHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	t, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "term not found"); return }
	response.Success(w, t, "")
}

// PUT /terms/{id} — no schoolID from context, correct as-is
func (h *TermHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.UpdateTermDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	if err := h.svc.Update(id, dto); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "term updated")
}

// POST /terms/{id}/set-current — service SetCurrent(schoolID, termID int64)
// nil guard + dereference required
func (h *TermHandler) SetCurrent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	if err := h.svc.SetCurrent(*schoolID, id); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "current term updated")
}

// DELETE /terms/{id} — no schoolID from context, correct as-is
func (h *TermHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "term deleted")
}