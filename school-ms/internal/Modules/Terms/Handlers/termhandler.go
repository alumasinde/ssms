package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Terms/DTOs"
	svc  "school-ms/internal/Modules/Terms/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type TermHandler struct{ svc *svc.TermService }
func NewTermHandler(s *svc.TermService) *TermHandler { return &TermHandler{svc: s} }

func (h *TermHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateTermDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	t, err := h.svc.Create(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, t, "term created")
}
func (h *TermHandler) ListByYear(w http.ResponseWriter, r *http.Request) {
	yearID, _ := strconv.ParseInt(chi.URLParam(r, "yearId"), 10, 64)
	list, err := h.svc.ListByYear(yearID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}
func (h *TermHandler) ListBySchool(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.ListBySchool(mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}
func (h *TermHandler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	t, err := h.svc.GetCurrent(mw.GetSchoolID(r.Context()))
	if err != nil { response.NotFound(w, "no active term found"); return }
	response.Success(w, t, "")
}
func (h *TermHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	t, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "term not found"); return }
	response.Success(w, t, "")
}
func (h *TermHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var dto dtos.UpdateTermDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	if err := h.svc.Update(id, dto); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "term updated")
}
func (h *TermHandler) SetCurrent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.SetCurrent(mw.GetSchoolID(r.Context()), id); err != nil {
		response.InternalError(w, err.Error()); return
	}
	response.Success(w, nil, "current term updated")
}
func (h *TermHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "term deleted")
}
