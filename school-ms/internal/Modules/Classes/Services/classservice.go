package services

import (
	dtos "school-ms/internal/Modules/Classes/DTOs"
	models "school-ms/internal/Modules/Classes/Models"
	repos "school-ms/internal/Modules/Classes/Repositories"
)

type ClassService struct{ repo *repos.ClassRepository }

func NewClassService(repo *repos.ClassRepository) *ClassService { return &ClassService{repo: repo} }

func (s *ClassService) Create(dto dtos.CreateClassDTO) (*models.Class, error) {
	c := &models.Class{SchoolID: dto.SchoolID, Name: dto.Name, Level: dto.Level, Stream: dto.Stream}
	return c, s.repo.Create(c)
}
func (s *ClassService) GetByID(id int64) (*models.Class, error) { return s.repo.FindByID(id) }
func (s *ClassService) List(schoolID int64) ([]models.Class, error) { return s.repo.ListBySchool(schoolID) }
func (s *ClassService) Update(id int64, dto dtos.CreateClassDTO) error {
	c := &models.Class{ID: id, SchoolID: dto.SchoolID, Name: dto.Name, Level: dto.Level, Stream: dto.Stream}
	return s.repo.Update(c)
}
func (s *ClassService) Delete(id int64) error { return s.repo.Delete(id) }
