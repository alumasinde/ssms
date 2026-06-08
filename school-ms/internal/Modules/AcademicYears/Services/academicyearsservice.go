package services

import (
	dtos "school-ms/internal/Modules/AcademicYears/DTOs"
	models "school-ms/internal/Modules/AcademicYears/Models"
	repos "school-ms/internal/Modules/AcademicYears/Repositories"
)

type AcademicYearService struct{ repo *repos.AcademicYearRepository }

func NewAcademicYearService(repo *repos.AcademicYearRepository) *AcademicYearService {
	return &AcademicYearService{repo: repo}
}

func (s *AcademicYearService) Create(dto dtos.CreateAcademicYearDTO) (*models.AcademicYear, error) {
	ay := &models.AcademicYear{
		SchoolID: dto.SchoolID, Name: dto.Name,
		StartDate: dto.StartDate, EndDate: dto.EndDate, IsCurrent: dto.IsCurrent,
	}
	return ay, s.repo.Create(ay)
}

func (s *AcademicYearService) GetByID(id int64) (*models.AcademicYear, error) {
	return s.repo.FindByID(id)
}

func (s *AcademicYearService) List(schoolID int64) ([]models.AcademicYear, error) {
	return s.repo.ListBySchool(schoolID)
}

func (s *AcademicYearService) GetCurrent(schoolID int64) (*models.AcademicYear, error) {
	return s.repo.GetCurrent(schoolID)
}

func (s *AcademicYearService) SetCurrent(schoolID, yearID int64) error {
	return s.repo.SetCurrent(schoolID, yearID)
}

func (s *AcademicYearService) Update(id int64, dto dtos.CreateAcademicYearDTO) error {
	ay := &models.AcademicYear{
		ID: id, SchoolID: dto.SchoolID, Name: dto.Name,
		StartDate: dto.StartDate, EndDate: dto.EndDate, IsCurrent: dto.IsCurrent,
	}
	return s.repo.Update(ay)
}

func (s *AcademicYearService) Delete(id int64) error {
	return s.repo.Delete(id)
}
