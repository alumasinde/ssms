package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Finance/DTOs"
	services "school-ms/internal/Modules/Finance/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type FinanceHandler struct{ svc *services.FinanceService }

func NewFinanceHandler(svc *services.FinanceService) *FinanceHandler { return &FinanceHandler{svc: svc} }

func (h *FinanceHandler) CreateFeeType(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateFeeTypeDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	ft, err := h.svc.CreateFeeType(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, ft, "fee type created")
}

func (h *FinanceHandler) ListFeeTypes(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.ListFeeTypes(mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *FinanceHandler) GenerateInvoices(w http.ResponseWriter, r *http.Request) {
	var dto dtos.GenerateInvoicesDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	count, err := h.svc.GenerateInvoices(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, map[string]int{"invoices_created": count}, "invoices generated")
}

func (h *FinanceHandler) StudentStatement(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	stmt, err := h.svc.GetStudentStatement(studentID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, stmt, "")
}

func (h *FinanceHandler) RecordPayment(w http.ResponseWriter, r *http.Request) {
	var dto dtos.RecordPaymentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	if err := h.svc.RecordPayment(dto); err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, nil, "payment recorded")
}

func (h *FinanceHandler) InvoicePayments(w http.ResponseWriter, r *http.Request) {
	invoiceID, _ := strconv.ParseInt(chi.URLParam(r, "invoiceId"), 10, 64)
	list, err := h.svc.GetInvoicePayments(invoiceID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}
