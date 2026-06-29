package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw   "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Exams/DTOs"
	svc  "school-ms/internal/Modules/Exams/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ExamHandler struct{ svc *svc.ExamService }

func NewExamHandler(s *svc.ExamService) *ExamHandler { return &ExamHandler{svc: s} }

func (h *ExamHandler) CreateExam(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateExamDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	// Service only takes dto — set SchoolID on the DTO directly
	dto.SchoolID = *schoolID
	e, err := h.svc.CreateExam(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, e, "exam created")
}

func (h *ExamHandler) ListExams(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if termID, _ := strconv.ParseInt(r.URL.Query().Get("term_id"), 10, 64); termID > 0 {
		list, err := h.svc.ListByTerm(termID)
		if err != nil { response.ServerError(w, err); return }
		response.Success(w, list, ""); return
	}
	if classID, _ := strconv.ParseInt(r.URL.Query().Get("class_id"), 10, 64); classID > 0 {
		list, err := h.svc.ListByClass(classID)
		if err != nil { response.ServerError(w, err); return }
		response.Success(w, list, ""); return
	}
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	list, err := h.svc.ListExams(*schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

func (h *ExamHandler) GetExam(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	e, err := h.svc.GetExam(id)
	if err != nil { response.NotFound(w, "exam not found"); return }
	response.Success(w, e, "")
}

func (h *ExamHandler) SubmitResults(w http.ResponseWriter, r *http.Request) {
	var dto dtos.SubmitResultDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	if err := h.svc.SubmitResults(dto, mw.GetUserID(r.Context()), *schoolID); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "results submitted")
}

func (h *ExamHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	examID, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if classID, _ := strconv.ParseInt(r.URL.Query().Get("class_id"), 10, 64); classID > 0 {
		list, err := h.svc.GetResultsByExamAndClass(examID, classID)
		if err != nil { response.ServerError(w, err); return }
		response.Success(w, list, ""); return
	}
	list, err := h.svc.GetResultsByExamEnriched(examID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

func (h *ExamHandler) GetStudentResults(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	list, err := h.svc.GetStudentResults(studentID, examID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

func (h *ExamHandler) CreateGradeScale(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateGradeScaleDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Forbidden(w, "no school context"); return
	}
	dto.SchoolID = *schoolID
	gs, err := h.svc.CreateGradeScale(dto)
	if err != nil { response.ServerError(w, err); return }
	response.Created(w, gs, "grade scale created")
}

func (h *ExamHandler) GetGradeScales(w http.ResponseWriter, r *http.Request) {
	schoolID := mw.GetSchoolID(r.Context())
	if schoolID == nil {
		response.Success(w, []interface{}{}, ""); return
	}
	list, err := h.svc.GetGradeScales(*schoolID)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, list, "")
}

func (h *ExamHandler) UpdateGradeScale(w http.ResponseWriter, r *http.Request) {
	var dto dtos.UpdateGradeScaleDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.BadRequest(w, "invalid payload"); return
	}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	gs, err := h.svc.UpdateGradeScale(id, dto)
	if err != nil { response.ServerError(w, err); return }
	response.Success(w, gs, "grade scale updated")
}

func (h *ExamHandler) DeleteGradeScale(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err := h.svc.DeleteGradeScale(id); err != nil {
		response.ServerError(w, err); return
	}
	response.Success(w, nil, "grade scale deleted")
}