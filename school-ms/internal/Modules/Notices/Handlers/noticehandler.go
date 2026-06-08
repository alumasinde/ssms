package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Notices/DTOs"
	services "school-ms/internal/Modules/Notices/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type NoticeHandler struct{ svc *services.NoticeService }

func NewNoticeHandler(svc *services.NoticeService) *NoticeHandler { return &NoticeHandler{svc: svc} }

func (h *NoticeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateNoticeDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	n, err := h.svc.Create(dto, mw.GetUserID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, n, "notice published")
}

func (h *NoticeHandler) List(w http.ResponseWriter, r *http.Request) {
	audience := r.URL.Query().Get("audience")
	list, err := h.svc.List(mw.GetSchoolID(r.Context()), audience)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *NoticeHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	n, err := h.svc.GetByID(id)
	if err != nil { response.NotFound(w, "notice not found"); return }
	response.Success(w, n, "")
}

func (h *NoticeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.Delete(id); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "notice deleted")
}
