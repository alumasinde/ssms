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

// Sentinel errors — exported so the handler uses errors.Is() cleanly.
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountDisabled    = errors.New("account is disabled")
)

type AuthService struct{ repo *repos.AuthRepository }

func NewAuthService(repo *repos.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Login validates credentials and returns a JWT pair plus contextual metadata.
func (s *AuthService) Login(req dtos.LoginRequest, tenantID int64) (*dtos.LoginResponse, error) {
	// 1. Fetch user — return identical error for "not found" and "wrong password"
	//    to prevent user-enumeration attacks.
	user, err := s.repo.FindUserByEmailAndTenant(req.Email, tenantID)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// 2. Active check — separate sentinel so the handler can give a clear 403.
	if !user.IsActive {
		return nil, ErrAccountDisabled
	}

	// 3. Password verification — constant-time compare via bcrypt.
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash), []byte(req.Password),
	); err != nil {
		return nil, ErrInvalidCredentials
	}

	// 4. Roles via user_roles → roles join (schema-aligned, two-table join).
	//    Non-fatal: a user with no roles can still log in but will have no access.
	roles, _ := s.repo.FindRoleNamesByUserID(user.ID)

	// 5. Permissions for those roles via role_permissions → permissions.
	//    Non-fatal for the same reason.
	perms, _ := s.repo.FindPermissionsByRoleNames(roles)

	// 6. Academic context — best-effort, never blocks login.
	var academicYearID, termID *int64
	if user.SchoolID != nil {
		if ayID, err := s.repo.GetCurrentAcademicYearID(*user.SchoolID); err == nil {
			academicYearID = ayID
			if tID, err := s.repo.GetCurrentTermID(*ayID); err == nil {
				termID = tID
			}
		}
	}

	// 7. Stamp last_login_at — fire-and-forget so it never blocks the response.
	go s.repo.UpdateLastLogin(user.ID)

	// 8. Sign token pair.
	//    Access token carries roles + academic context for zero-DB-query auth.
	//    Refresh token is kept slim (no context — re-populated on next login).
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

// Refresh validates a refresh token and issues a new access + refresh pair.
// Permissions are re-fetched from DB so role changes take effect immediately.
func (s *AuthService) Refresh(refreshToken string) (*dtos.LoginResponse, error) {
	claims, err := s.parseToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	// Re-fetch permissions in case roles changed since the token was issued
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

// signToken builds and signs a JWT.
// All pointer fields (*int64) match the Claims struct in middleware/auth.go exactly.
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

// parseToken validates a JWT string and extracts its claims.
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
