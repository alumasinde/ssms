package services

import (
	"errors"
	"time"

	"school-ms/config"
	mw "school-ms/internal/middleware"
	dtos "school-ms/internal/Modules/Auth/DTOs"
	repos "school-ms/internal/Modules/Auth/Repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repos.AuthRepository
}

func NewAuthService(repo *repos.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Login resolves the user by email within the already-resolved tenantID,
// verifies the password, and returns a token pair + user summary.
func (s *AuthService) Login(req dtos.LoginRequest, tenantID int64) (*dtos.LoginResponse, error) {
	if tenantID == 0 {
		return nil, errors.New("tenant could not be resolved from domain — please contact support")
	}

	user, err := s.repo.FindUserByEmailAndTenant(req.Email, tenantID)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	access, err := s.sign(user.ID, user.TenantID, user.SchoolID, user.Role, config.App.JWTAccessTTL)
	if err != nil {
		return nil, err
	}
	refresh, err := s.sign(user.ID, user.TenantID, user.SchoolID, user.Role, config.App.JWTRefreshTTL)
	if err != nil {
		return nil, err
	}

	perms, err := s.repo.FindUserPermissions(user.Role)
	if err != nil {
		perms = []string{}
	}

	return &dtos.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		Permissions:  perms,
		User: dtos.UserSummary{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Role:     user.Role,
			TenantID: user.TenantID,
			SchoolID: user.SchoolID,
		},
	}, nil
}

func (s *AuthService) sign(userID, tenantID, schoolID int64, role string, ttl time.Duration) (string, error) {
	claims := &mw.Claims{
		UserID:   userID,
		TenantID: tenantID,
		SchoolID: schoolID,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(config.App.JWTSecret))
}
