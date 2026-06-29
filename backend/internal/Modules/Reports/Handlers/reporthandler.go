package handlers

import (
	"net/http"
	"strconv"

	mw       "school-ms/internal/middleware"
	services "school-ms/internal/Modules/Reports/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ReportHandler struct{ svc *services.ReportService }

func NewReportHandler(svc *services.ReportService) *ReportHandler { return &ReportHandler{svc: svc} }

// GET /reports/report-card/{studentId}?exam_id=
// no schoolID involved — correct as-is
func (h *ReportHandler) ReportCard(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	card, err := h.svc.GetReportCard(studentID, examID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, card, "")
}

// GET /reports/class-results?exam_id=
// no schoolID involved — correct as-is
func (h *ReportHandler) ClassResults(w http.ResponseWriter, r *http.Request) {
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	rows, err := h.svc.GetClassResults(examID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, rows, "")
}

// GET /reports/fee-collection?term_id=
// service GetFeeCollection(schoolID int64, termID int64)
// nil guard + dereference required
func (h *ReportHandler) FeeCollection(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	report, err := h.svc.GetFeeCollection(*schoolID, termID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, report, "")
}

// GET /reports/attendance-summary?term_id=
// service GetAttendanceSummary(schoolID int64, termID int64)
// nil guard + dereference required
func (h *ReportHandler) AttendanceSummary(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64)
	reports, err := h.svc.GetAttendanceSummary(*schoolID, termID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, reports, "")
}

// GET /reports/subject-performance?exam_id=
// no schoolID involved — correct as-is
func (h *ReportHandler) SubjectPerformance(w http.ResponseWriter, r *http.Request) {
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	rows, err := h.svc.GetSubjectPerformance(examID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, rows, "")
}
