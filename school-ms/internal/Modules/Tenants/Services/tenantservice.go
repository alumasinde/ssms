package services

import (
	dtos "school-ms/internal/Modules/Tenants/DTOs"
	models "school-ms/internal/Modules/Tenants/Models"
	repos "school-ms/internal/Modules/Tenants/Repositories"
)

type TenantService struct {
	repo *repos.TenantRepository
}

func NewTenantService(repo *repos.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (s *TenantService) Create(dto dtos.CreateTenantDTO) (*models.Tenant, error) {
	t := &models.Tenant{
		Slug: dto.Slug, Name: dto.Name, Domain: dto.Domain, Plan: dto.Plan,
	}
	return t, s.repo.Create(t)
}

func (s *TenantService) GetByID(id int64) (*models.Tenant, error) {
	return s.repo.FindByID(id)
}

func (s *TenantService) GetBySlug(slug string) (*models.Tenant, error) {
	return s.repo.FindBySlug(slug)
}

func (s *TenantService) List() ([]models.Tenant, error) {
	return s.repo.List()
}

func (s *TenantService) Update(id int64, dto dtos.UpdateTenantDTO) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	name, domain, plan, isActive := existing.Name, existing.Domain, existing.Plan, existing.IsActive
	if dto.Name != "" {
		name = dto.Name
	}
	if dto.Domain != "" {
		domain = dto.Domain
	}
	if dto.Plan != "" {
		plan = dto.Plan
	}
	if dto.IsActive != nil {
		isActive = *dto.IsActive
	}
	return s.repo.Update(id, name, domain, plan, isActive)
}

func (s *TenantService) Delete(id int64) error {
	return s.repo.Delete(id)
}
