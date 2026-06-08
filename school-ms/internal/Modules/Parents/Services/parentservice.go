package services

import (
	dtos "school-ms/internal/Modules/Parents/DTOs"
	models "school-ms/internal/Modules/Parents/Models"
	repos "school-ms/internal/Modules/Parents/Repositories"
)

type ParentService struct{ repo *repos.ParentRepository }

func NewParentService(repo *repos.ParentRepository) *ParentService { return &ParentService{repo: repo} }

func (s *ParentService) Create(dto dtos.CreateParentDTO) (*models.Parent, error) {
	p := &models.Parent{UserID: dto.UserID, SchoolID: dto.SchoolID, Phone: dto.Phone, Occupation: dto.Occupation, Address: dto.Address}
	return p, s.repo.Create(p)
}
func (s *ParentService) GetByID(id int64) (*models.ParentDetail, error) { return s.repo.FindByID(id) }
func (s *ParentService) List(schoolID int64) ([]models.ParentDetail, error) { return s.repo.ListBySchool(schoolID) }
func (s *ParentService) LinkStudent(dto dtos.LinkStudentDTO) error {
	return s.repo.LinkStudent(&models.ParentStudent{ParentID: dto.ParentID, StudentID: dto.StudentID, Relationship: dto.Relationship})
}
func (s *ParentService) GetStudentParents(studentID int64) ([]models.ParentDetail, error) {
	return s.repo.GetParentsByStudent(studentID)
}
