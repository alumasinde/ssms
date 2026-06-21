package services

import (
	dtos   "school-ms/internal/Modules/Terms/DTOs"
	models "school-ms/internal/Modules/Terms/Models"
	repos  "school-ms/internal/Modules/Terms/Repositories"
)

type TermService struct{ repo *repos.TermRepository }

func NewTermService(repo *repos.TermRepository) *TermService { return &TermService{repo: repo} }

func (s *TermService) Create(dto dtos.CreateTermDTO) (*models.Term, error) {
	t := &models.Term{AcademicYearID: dto.AcademicYearID, Name: dto.Name,
		StartDate: dto.StartDate, EndDate: dto.EndDate, IsCurrent: dto.IsCurrent}
	return t, s.repo.Create(t)
}
func (s *TermService) GetByID(id int64) (*models.Term, error)             { return s.repo.FindByID(id) }
func (s *TermService) ListByYear(yearID int64) ([]models.Term, error)     { return s.repo.ListByAcademicYear(yearID) }
func (s *TermService) ListBySchool(schoolID int64) ([]models.Term, error) { return s.repo.ListBySchool(schoolID) }
func (s *TermService) GetCurrent(schoolID int64) (*models.Term, error)    { return s.repo.GetCurrent(schoolID) }
func (s *TermService) SetCurrent(schoolID, termID int64) error            { return s.repo.SetCurrent(schoolID, termID) }
func (s *TermService) Delete(id int64) error                              { return s.repo.Delete(id) }
func (s *TermService) Update(id int64, dto dtos.UpdateTermDTO) error {
	return s.repo.Update(&models.Term{ID: id, Name: dto.Name, StartDate: dto.StartDate, EndDate: dto.EndDate})
}
