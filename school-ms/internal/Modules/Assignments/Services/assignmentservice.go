package services

import (
	dtos   "school-ms/internal/Modules/Assignments/DTOs"
	models "school-ms/internal/Modules/Assignments/Models"
	repos  "school-ms/internal/Modules/Assignments/Repositories"
)

type AssignmentService struct{ repo *repos.AssignmentRepository }
func NewAssignmentService(r *repos.AssignmentRepository) *AssignmentService { return &AssignmentService{repo:r} }

func (s *AssignmentService) Create(dto dtos.CreateAssignmentDTO, schoolID int64) (*models.Assignment,error) {
	a := &models.Assignment{SchoolID:schoolID,ClassID:dto.ClassID,SubjectID:dto.SubjectID,
		TeacherID:dto.TeacherID,TermID:dto.TermID,Title:dto.Title,
		Description:dto.Description,DueDate:dto.DueDate,MaxMarks:dto.MaxMarks}
	return a, s.repo.Create(a)
}
func (s *AssignmentService) ListByClass(classID,termID int64) ([]models.AssignmentDetail,error) { return s.repo.ListByClass(classID,termID) }
func (s *AssignmentService) ListByTeacher(teacherID,termID int64) ([]models.AssignmentDetail,error) { return s.repo.ListByTeacher(teacherID,termID) }
func (s *AssignmentService) ListBySchool(schoolID,termID int64) ([]models.AssignmentDetail,error) { return s.repo.ListBySchool(schoolID,termID) }
func (s *AssignmentService) Delete(id int64) error { return s.repo.Delete(id) }
