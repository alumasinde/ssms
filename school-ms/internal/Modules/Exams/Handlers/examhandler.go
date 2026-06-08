package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Exams/DTOs"
	services "school-ms/internal/Modules/Exams/Services"
	"school-ms/internal/pkg/response"

	"github.com/go-chi/chi/v5"
)

type ExamHandler struct{ svc *services.ExamService }

func NewExamHandler(svc *services.ExamService) *ExamHandler { return &ExamHandler{svc: svc} }

func (h *ExamHandler) CreateExam(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateExamDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	e, err := h.svc.CreateExam(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, e, "exam created")
}

func (h *ExamHandler) ListExams(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.ListExams(mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
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
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	if err := h.svc.SubmitResults(dto, mw.GetUserID(r.Context()), mw.GetSchoolID(r.Context())); err != nil {
		response.InternalError(w, err.Error()); return
	}
	response.Success(w, nil, "results submitted")
}

func (h *ExamHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	examID, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	list, err := h.svc.GetResultsByExam(examID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *ExamHandler) GetStudentResults(w http.ResponseWriter, r *http.Request) {
	studentID, _ := strconv.ParseInt(chi.URLParam(r, "studentId"), 10, 64)
	examID, _ := strconv.ParseInt(r.URL.Query().Get("exam_id"), 10, 64)
	list, err := h.svc.GetStudentResults(studentID, examID)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}

func (h *ExamHandler) CreateGradeScale(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateGradeScaleDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil { response.BadRequest(w, "invalid payload"); return }
	dto.SchoolID = mw.GetSchoolID(r.Context())
	gs, err := h.svc.CreateGradeScale(dto)
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Created(w, gs, "grade scale created")
}

func (h *ExamHandler) GetGradeScales(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.GetGradeScales(mw.GetSchoolID(r.Context()))
	if err != nil { response.InternalError(w, err.Error()); return }
	response.Success(w, list, "")
}
