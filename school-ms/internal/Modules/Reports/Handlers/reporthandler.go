package handlers

import (
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	services "school-ms/internal/Modules/Reports/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ReportHandler struct{ svc *services.ReportService }

func NewReportHandler(svc *services.ReportService) *ReportHandler { return &ReportHandler{svc: svc} }

func (h *ReportHandler) ReportCard(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	card, err := h.svc.GetReportCard(studentID, examID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, card, "")
}

func (h *ReportHandler) ClassResults(w http.ResponseWriter, r *http.Request) {
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	rows, err := h.svc.GetClassResults(examID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, rows, "")
}

func (h *ReportHandler) FeeCollection(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	report, err := h.svc.GetFeeCollection(schoolID, termID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, report, "")
}

func (h *ReportHandler) AttendanceSummary(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	reports, err := h.svc.GetAttendanceSummary(schoolID, termID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, reports, "")
}

func (h *ReportHandler) SubjectPerformance(w http.ResponseWriter, r *http.Request) {
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	rows, err := h.svc.GetSubjectPerformance(examID)
	if err != nil {
		response.InternalError(w, err.Error())
		return
	}
	response.Success(w, rows, "")
}
