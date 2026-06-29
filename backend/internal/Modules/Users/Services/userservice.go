package services
import (
	"errors"
	"strings"

	dtos   "school-ms/internal/Modules/Users/DTOs"
	models "school-ms/internal/Modules/Users/Models"
	repos  "school-ms/internal/Modules/Users/Repositories"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailTaken    = errors.New("email already registered for this tenant")
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("current password is incorrect")
	ErrInvalidRole   = errors.New("role_code is invalid or inactive")
)

type UserService struct{ repo *repos.UserRepository }

func NewUserService(repo *repos.UserRepository) *UserService { return &UserService{repo: repo} }

// Register creates a user row then assigns the role via user_roles.
func (s *UserService) Register(dto dtos.RegisterDTO, tenantID int64, schoolID *int64) (*models.UserWithRoles, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil { return nil, err }
	u := &models.User{
		TenantID: tenantID, SchoolID: schoolID,
		FirstName: dto.FirstName, LastName: dto.LastName,
		Email: dto.Email, PasswordHash: string(hash), Phone: dto.Phone,
	}
	if err := s.repo.Create(u); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") ||
			strings.Contains(err.Error(), "uq_users_email_tenant") {
			return nil, ErrEmailTaken
		}
		return nil, err
	}
	if err := s.repo.AssignRole(u.ID, tenantID, dto.RoleCode); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, ErrInvalidRole
		}
		return nil, err
	}
	roles, _ := s.repo.FindRoleNames(u.ID)
	return &models.UserWithRoles{User: *u, Roles: roles}, nil
}

func (s *UserService) GetByID(id int64) (*models.UserWithRoles, error) {
	u, err := s.repo.FindByID(id)
	if err != nil { return nil, ErrUserNotFound }
	roles, _ := s.repo.FindRoleNames(u.ID)
	return &models.UserWithRoles{User: *u, Roles: roles}, nil
}

func (s *UserService) ListByTenant(tenantID int64) ([]models.UserWithRoles, error) {
	users, err := s.repo.ListByTenant(tenantID)
	if err != nil { return nil, err }
	return s.enrich(users), nil
}

func (s *UserService) ListBySchool(schoolID int64) ([]models.UserWithRoles, error) {
	users, err := s.repo.ListBySchool(schoolID)
	if err != nil { return nil, err }
	return s.enrich(users), nil
}

func (s *UserService) ChangePassword(userID int64, dto dtos.ChangePasswordDTO) error {
	u, err := s.repo.FindByID(userID)
	if err != nil { return ErrUserNotFound }
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(dto.OldPassword)); err != nil {
		return ErrWrongPassword
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
	if err != nil { return err }
	return s.repo.UpdatePassword(userID, string(hash))
}

func (s *UserService) UpdateRole(id, tenantID int64, dto dtos.UpdateRoleDTO) error {
	return s.repo.ReplaceRole(id, tenantID, dto.RoleCode)
}

func (s *UserService) Deactivate(id, actorID int64) error {
	return s.repo.SoftDelete(id, actorID)
}

func (s *UserService) Activate(id int64) error { return s.repo.Activate(id) }

func (s *UserService) enrich(users []models.User) []models.UserWithRoles {
	out := make([]models.UserWithRoles, 0, len(users))
	for _, u := range users {
		roles, _ := s.repo.FindRoleNames(u.ID)
		out = append(out, models.UserWithRoles{User: u, Roles: roles})
	}
	return out
}