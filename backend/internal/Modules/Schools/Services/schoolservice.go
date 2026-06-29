package services

import (
	dtos "school-ms/internal/Modules/Schools/DTOs"
	models "school-ms/internal/Modules/Schools/Models"
	repos "school-ms/internal/Modules/Schools/Repositories"
)

type SchoolService struct {
	repo *repos.SchoolRepository
}

func NewSchoolService(repo *repos.SchoolRepository) *SchoolService {
	return &SchoolService{repo: repo}
}

func (s *SchoolService) CreateSchool(dto dtos.CreateSchoolDTO) (*models.School, error) {
	school := &models.School{
		TenantID: dto.TenantID,
		Name:     dto.Name,
		Code:     dto.Code,
		Address:  dto.Address,
		Phone:    dto.Phone,
		Email:    dto.Email,
		LogoURL:  dto.LogoURL,
	}
	return school, s.repo.CreateSchool(school)
}

func (s *SchoolService) GetSchool(id int64) (*models.School, error) {
	return s.repo.FindSchoolByID(id)
}

func (s *SchoolService) ListSchools(tenantID int64) ([]models.School, error) {
	return s.repo.ListSchoolsByTenant(tenantID)
}

func (s *SchoolService) CreateAcademicYear(dto dtos.CreateAcademicYearDTO) (*models.AcademicYear, error) {
	ay := &models.AcademicYear{
		SchoolID:  dto.SchoolID,
		Name:      dto.Name,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
		IsCurrent: dto.IsCurrent,
	}
	return ay, s.repo.CreateAcademicYear(ay)
}

func (s *SchoolService) ListAcademicYears(schoolID int64) ([]models.AcademicYear, error) {
	return s.repo.ListAcademicYears(schoolID)
}

func (s *SchoolService) GetCurrentAcademicYear(schoolID int64) (*models.AcademicYear, error) {
	return s.repo.GetCurrentAcademicYear(schoolID)
}

func (s *SchoolService) CreateTerm(dto dtos.CreateTermDTO) (*models.Term, error) {
	t := &models.Term{
		AcademicYearID: dto.AcademicYearID,
		Name:           dto.Name,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		IsCurrent:      dto.IsCurrent,
	}
	return t, s.repo.CreateTerm(t)
}

func (s *SchoolService) ListTerms(academicYearID int64) ([]models.Term, error) {
	return s.repo.ListTerms(academicYearID)
}

func (s *SchoolService) GetCurrentTerm(schoolID int64) (*models.Term, error) {
	return s.repo.GetCurrentTerm(schoolID)
}

func (s *SchoolService) UpdateSchool(id int64, dto dtos.CreateSchoolDTO) error {
	school, err := s.repo.FindSchoolByID(id)
	if err != nil {
		return err
	}
	if dto.Name != "" { school.Name = dto.Name }
	if dto.Code != "" { school.Code = dto.Code }
	if dto.Address != "" { school.Address = dto.Address }
	if dto.Phone != "" { school.Phone = dto.Phone }
	if dto.Email != "" { school.Email = dto.Email }
	if dto.LogoURL != "" { school.LogoURL = dto.LogoURL }
	return s.repo.UpdateSchool(school)
}
