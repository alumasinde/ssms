package services

import (
	"errors"
	"time"

	"school-ms/config"
	mw "school-ms/internal/middleware"

	dtos  "school-ms/internal/Modules/Auth/DTOs"
	repos "school-ms/internal/Modules/Auth/Repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountDisabled    = errors.New("account is disabled")
	ErrTokenBlacklisted   = errors.New("token has been revoked")
)

type AuthService struct{ repo *repos.AuthRepository }

func NewAuthService(repo *repos.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(req dtos.LoginRequest, tenantID int64) (*dtos.LoginResponse, error) {
	user, err := s.repo.FindUserByEmailAndTenant(req.Email, tenantID)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if !user.IsActive {
		return nil, ErrAccountDisabled
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash), []byte(req.Password),
	); err != nil {
		return nil, ErrInvalidCredentials
	}

	roles, _ := s.repo.FindRoleNamesByUserID(user.ID)
	perms, _ := s.repo.FindPermissionsByRoleNames(roles)

	var academicYearID, termID *int64
	if user.SchoolID != nil {
		if ayID, err := s.repo.GetCurrentAcademicYearID(*user.SchoolID); err == nil {
			academicYearID = ayID
			if tID, err := s.repo.GetCurrentTermID(*ayID); err == nil {
				termID = tID
			}
		}
	}

	go s.repo.UpdateLastLogin(user.ID)

	access, err := s.signToken(
		user.ID, user.TenantID, user.SchoolID,
		roles, academicYearID, termID,
		config.App.JWTAccessTTL,
	)
	if err != nil {
		return nil, errors.New("could not issue access token")
	}

	refresh, err := s.signToken(
		user.ID, user.TenantID, user.SchoolID,
		roles, nil, nil,
		config.App.JWTRefreshTTL,
	)
	if err != nil {
		return nil, errors.New("could not issue refresh token")
	}

	fullName := user.FullName()

	return &dtos.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    int64(config.App.JWTAccessTTL.Seconds()),
		Roles:        roles,
		Permissions:  perms,
		User: dtos.UserSummary{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Name:      fullName,
			Email:     user.Email,
			TenantID:  user.TenantID,
			SchoolID:  user.SchoolID,
			IsActive:  user.IsActive,
		},
		Context: dtos.AuthContext{
			AcademicYearID: academicYearID,
			TermID:         termID,
		},
	}, nil
}

func (s *AuthService) Refresh(refreshToken string) (*dtos.LoginResponse, error) {
	claims, err := s.parseToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	// FIX: check token blacklist before issuing new tokens
	if blacklisted, _ := s.repo.IsTokenBlacklisted(claims.ID); blacklisted {
		return nil, ErrTokenBlacklisted
	}

	perms, _ := s.repo.FindPermissionsByRoleNames(claims.Roles)

	access, err := s.signToken(
		claims.UserID, claims.TenantID, claims.SchoolID,
		claims.Roles,
		claims.AcademicYearID, claims.TermID,
		config.App.JWTAccessTTL,
	)
	if err != nil {
		return nil, errors.New("could not issue access token")
	}

	newRefresh, err := s.signToken(
		claims.UserID, claims.TenantID, claims.SchoolID,
		claims.Roles, nil, nil,
		config.App.JWTRefreshTTL,
	)
	if err != nil {
		return nil, errors.New("could not issue refresh token")
	}

	return &dtos.LoginResponse{
		AccessToken:  access,
		RefreshToken: newRefresh,
		ExpiresIn:    int64(config.App.JWTAccessTTL.Seconds()),
		Roles:        claims.Roles,
		Permissions:  perms,
	}, nil
}

// Logout blacklists the refresh token JTI so it cannot be reused.
// FIX: previous implementation was a no-op — tokens lived for 7 days
// after logout with no way to revoke them.
func (s *AuthService) Logout(refreshToken string, userID int64) error {
	claims, err := s.parseToken(refreshToken)
	if err != nil {
		// Token already invalid — nothing to blacklist
		return nil
	}
	expiresAt := claims.ExpiresAt.Time
	return s.repo.BlacklistToken(claims.ID, userID, expiresAt)
}

func (s *AuthService) signToken(
	userID, tenantID int64,
	schoolID *int64,
	roles []string,
	academicYearID, termID *int64,
	ttl time.Duration,
) (string, error) {
	claims := &mw.Claims{
		UserID:         userID,
		TenantID:       tenantID,
		SchoolID:       schoolID,
		Roles:          roles,
		AcademicYearID: academicYearID,
		TermID:         termID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(config.App.JWTSecret))
}

func (s *AuthService) parseToken(tokenStr string) (*mw.Claims, error) {
	claims := &mw.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.App.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
