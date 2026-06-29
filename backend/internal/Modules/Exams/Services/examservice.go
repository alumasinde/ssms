package services

import (
	"fmt"
	dtos   "school-ms/internal/Modules/Exams/DTOs"
	models "school-ms/internal/Modules/Exams/Models"
	repos  "school-ms/internal/Modules/Exams/Repositories"
)

type ExamService struct{ repo *repos.ExamRepository }

func NewExamService(repo *repos.ExamRepository) *ExamService { return &ExamService{repo: repo} }

func (s *ExamService) CreateExam(dto dtos.CreateExamDTO) (*models.Exam, error) {
	e := &models.Exam{
		SchoolID: dto.SchoolID, TermID: dto.TermID, ClassID: dto.ClassID,
		Name: dto.Name, Type: dto.Type, StartDate: dto.StartDate, EndDate: dto.EndDate,
	}
	return e, s.repo.CreateExam(e)
}

func (s *ExamService) ListExams(schoolID int64) ([]models.Exam, error)         { return s.repo.ListBySchool(schoolID) }
func (s *ExamService) ListByTerm(termID int64) ([]models.Exam, error)          { return s.repo.ListByTerm(termID) }
func (s *ExamService) ListByClass(classID int64) ([]models.Exam, error)        { return s.repo.ListByClass(classID) }
func (s *ExamService) GetExam(id int64) (*models.Exam, error)                  { return s.repo.FindByID(id) }
func (s *ExamService) GetGradeScales(schoolID int64) ([]models.GradeScale, error) { return s.repo.GetGradeScales(schoolID) }
func (s *ExamService) UpdateGradeScale(id int64, dto dtos.UpdateGradeScaleDTO) (*models.GradeScale, error) {
	gs := &models.GradeScale{
		Grade:     dto.Grade,
		MinScore:  dto.MinScore,
		MaxScore:  dto.MaxScore,
		Remark:    dto.Remark,
	}
	return gs, s.repo.UpdateGradeScale(id, gs)
}

func (s *ExamService) DeleteGradeScale(id int64) error {
	if _, err := s.repo.FindGradeScaleByID(id); err != nil {
		return err
	}

	return s.repo.DeleteGradeScale(id)
}

func (s *ExamService) SubmitResults(
    dto dtos.SubmitResultDTO,
    gradedBy,
    schoolID int64,
) error {

    if dto.ClassID <= 0 {
        return fmt.Errorf("please select a class before submitting")
    }

    scales, err := s.repo.LoadGradeScales(schoolID)
    if err != nil {
        return err
    }

    records := make([]models.ExamResult, 0, len(dto.Results))

    for _, entry := range dto.Results {
        records = append(records, models.ExamResult{
            ExamID:    dto.ExamID,
            StudentID: entry.StudentID,
            SubjectID: entry.SubjectID,
            ClassID:   dto.ClassID,
            GradedBy:  gradedBy,
            Marks:     entry.Marks,
            MaxMarks:  entry.MaxMarks,
            Remarks:   entry.Remarks,
        })
    }

    return s.repo.BulkUpsertResults(records, scales)
}

func (s *ExamService) GetStudentResults(studentID, examID int64) ([]models.ExamResult, error) {
	return s.repo.GetStudentResults(studentID, examID)
}

func (s *ExamService) CreateGradeScale(dto dtos.CreateGradeScaleDTO) (*models.GradeScale, error) {
	gs := &models.GradeScale{SchoolID: dto.SchoolID, Grade: dto.Grade, MinScore: dto.MinScore, MaxScore: dto.MaxScore, Remark: dto.Remark}
	return gs, s.repo.CreateGradeScale(gs)
}

func (s *ExamService) GetResultsByExamEnriched(examID int64) ([]models.ExamResultDetail, error) {
	return s.repo.GetResultsByExamEnriched(examID)
}

func (s *ExamService) GetResultsByExamAndClass(examID, classID int64) ([]models.ExamResultDetail, error) {
	return s.repo.GetResultsByExamAndClass(examID, classID)
}
