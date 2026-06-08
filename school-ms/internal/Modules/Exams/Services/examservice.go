package services

import (
	dtos "school-ms/internal/Modules/Exams/DTOs"
	models "school-ms/internal/Modules/Exams/Models"
	repos "school-ms/internal/Modules/Exams/Repositories"
)

type ExamService struct{ repo *repos.ExamRepository }

func NewExamService(repo *repos.ExamRepository) *ExamService { return &ExamService{repo: repo} }

func (s *ExamService) CreateExam(dto dtos.CreateExamDTO) (*models.Exam, error) {
	e := &models.Exam{SchoolID: dto.SchoolID, TermID: dto.TermID, Name: dto.Name, Type: dto.Type, StartDate: dto.StartDate, EndDate: dto.EndDate}
	return e, s.repo.CreateExam(e)
}

func (s *ExamService) ListExams(schoolID int64) ([]models.Exam, error) {
	return s.repo.ListBySchool(schoolID)
}

func (s *ExamService) GetExam(id int64) (*models.Exam, error) {
	return s.repo.FindByID(id)
}

func (s *ExamService) SubmitResults(dto dtos.SubmitResultDTO, gradedBy, schoolID int64) error {
	var records []models.ExamResult
	for _, entry := range dto.Results {
		percent := 0.0
		if entry.MaxMarks > 0 {
			percent = entry.Marks / entry.MaxMarks * 100
		}
		grade := s.repo.ResolveGrade(schoolID, percent)
		records = append(records, models.ExamResult{
			ExamID: dto.ExamID, StudentID: entry.StudentID, SubjectID: entry.SubjectID,
			GradedBy: gradedBy, Marks: entry.Marks, MaxMarks: entry.MaxMarks,
			Grade: grade, Remarks: entry.Remarks,
		})
	}
	return s.repo.BulkUpsertResults(records)
}

func (s *ExamService) GetResultsByExam(examID int64) ([]models.ExamResult, error) {
	return s.repo.GetResultsByExam(examID)
}

func (s *ExamService) GetStudentResults(studentID, examID int64) ([]models.ExamResult, error) {
	return s.repo.GetStudentResults(studentID, examID)
}

func (s *ExamService) GetGradeScales(schoolID int64) ([]models.GradeScale, error) {
	return s.repo.GetGradeScales(schoolID)
}

func (s *ExamService) CreateGradeScale(dto dtos.CreateGradeScaleDTO) (*models.GradeScale, error) {
	gs := &models.GradeScale{SchoolID: dto.SchoolID, Grade: dto.Grade, MinScore: dto.MinScore, MaxScore: dto.MaxScore, Remark: dto.Remark}
	return gs, s.repo.CreateGradeScale(gs)
}
