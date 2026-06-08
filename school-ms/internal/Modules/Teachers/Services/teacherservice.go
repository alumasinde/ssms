package services

import (
	dtos "school-ms/internal/Modules/Teachers/DTOs"
	models "school-ms/internal/Modules/Teachers/Models"
	repos "school-ms/internal/Modules/Teachers/Repositories"
)

type TeacherService struct{ repo *repos.TeacherRepository }

func NewTeacherService(repo *repos.TeacherRepository) *TeacherService {
	return &TeacherService{repo: repo}
}

func (s *TeacherService) Create(dto dtos.CreateTeacherDTO) (*models.Teacher, error) {
	t := &models.Teacher{
		UserID: dto.UserID, SchoolID: dto.SchoolID, EmployeeNo: dto.EmployeeNo,
		Phone: dto.Phone, Gender: dto.Gender, DOB: dto.DOB,
		Qualification: dto.Qualification, PhotoURL: dto.PhotoURL,
	}
	return t, s.repo.Create(t)
}

func (s *TeacherService) GetByID(id int64) (*models.TeacherDetail, error) { return s.repo.FindByID(id) }

func (s *TeacherService) List(schoolID int64) ([]models.TeacherDetail, error) {
	return s.repo.ListBySchool(schoolID)
}

func (s *TeacherService) AssignSubject(dto dtos.AssignSubjectDTO) error {
	return s.repo.AssignSubject(&models.TeacherSubject{
		TeacherID: dto.TeacherID, SubjectID: dto.SubjectID, ClassID: dto.ClassID,
	})
}

func (s *TeacherService) GetSubjects(teacherID int64) ([]models.TeacherSubject, error) {
	return s.repo.GetSubjectsByTeacher(teacherID)
}
