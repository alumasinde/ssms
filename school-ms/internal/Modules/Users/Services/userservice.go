package services

import (
	"errors"

	dtos "school-ms/internal/Modules/Users/DTOs"
	models "school-ms/internal/Modules/Users/Models"
	repos "school-ms/internal/Modules/Users/Repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repos.UserRepository
}

func NewUserService(repo *repos.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(dto dtos.RegisterDTO) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &models.User{
		TenantID:     dto.TenantID,
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: string(hash),
		Role:         dto.Role,
	}
	return u, s.repo.Create(u)
}

func (s *UserService) GetByID(id int64) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) ListByTenant(tenantID int64) ([]models.User, error) {
	return s.repo.ListByTenant(tenantID)
}

func (s *UserService) ChangePassword(userID int64, dto dtos.ChangePasswordDTO) error {
	u, err := s.repo.FindByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(dto.OldPassword)); err != nil {
		return errors.New("old password incorrect")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.repo.UpdatePassword(userID, string(hash))
}

func (s *UserService) Deactivate(id int64) error {
	return s.repo.Deactivate(id)
}
