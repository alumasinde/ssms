package services

import (
	dtos "school-ms/internal/Modules/Subjects/DTOs"
	models "school-ms/internal/Modules/Subjects/Models"
	repos "school-ms/internal/Modules/Subjects/Repositories"
)

type SubjectService struct{ repo *repos.SubjectRepository }

func NewSubjectService(repo *repos.SubjectRepository) *SubjectService { return &SubjectService{repo: repo} }

func (s *SubjectService) Create(dto dtos.CreateSubjectDTO) (*models.Subject, error) {
	sub := &models.Subject{SchoolID: dto.SchoolID, Name: dto.Name, Code: dto.Code}
	return sub, s.repo.Create(sub)
}
func (s *SubjectService) List(schoolID int64) ([]models.Subject, error) { return s.repo.ListBySchool(schoolID) }
func (s *SubjectService) GetByID(id int64) (*models.Subject, error)     { return s.repo.FindByID(id) }
func (s *SubjectService) Update(id int64, dto dtos.CreateSubjectDTO) error {
	return s.repo.Update(&models.Subject{ID: id, Name: dto.Name, Code: dto.Code})
}
func (s *SubjectService) Delete(id int64) error { return s.repo.Delete(id) }
