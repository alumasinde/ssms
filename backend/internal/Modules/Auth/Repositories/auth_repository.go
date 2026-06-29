package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	models "school-ms/internal/Modules/Auth/Models"

	"github.com/jmoiron/sqlx"
)

const authTimeout = 5 * time.Second

type AuthRepository struct{ db *sqlx.DB }

func NewAuthRepository(db *sqlx.DB) *AuthRepository { return &AuthRepository{db: db} }

func (r *AuthRepository) FindUserByEmailAndTenant(
	email string, tenantID int64,
) (*models.LoginUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	email = strings.TrimSpace(strings.ToLower(email))

	var user models.LoginUser
	err := r.db.GetContext(ctx, &user, `
		SELECT
			u.id,
			u.tenant_id,
			u.school_id,
			u.first_name,
			u.last_name,
			u.email,
			u.password_hash,
			u.is_active
		FROM users u
		WHERE u.email     = ?
		  AND u.tenant_id = ?
		  AND u.is_active = 1
		  AND (u.deleted_at IS NULL OR u.deleted_at > NOW())
		LIMIT 1
	`, email, tenantID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) FindRoleNamesByUserID(userID int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	var roles []string
	err := r.db.SelectContext(ctx, &roles, `
		SELECT r.code
		FROM roles r
		INNER JOIN user_roles ur ON ur.role_id = r.id
		WHERE ur.user_id  = ?
		  AND r.is_active = 1
		ORDER BY r.code
	`, userID)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *AuthRepository) FindPermissionsByRoleNames(roleNames []string) ([]string, error) {
	if len(roleNames) == 0 {
		return []string{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	placeholders := make([]string, len(roleNames))
	args := make([]interface{}, len(roleNames))
	for i, code := range roleNames {
		placeholders[i] = "?"
		args[i] = code
	}

	query := fmt.Sprintf(`
		SELECT DISTINCT p.name
		FROM permissions p
		INNER JOIN role_permissions rp ON rp.permission_id = p.id
		INNER JOIN roles r             ON r.id             = rp.role_id
		WHERE r.code    IN (%s)
		  AND r.is_active = 1
		ORDER BY p.name
	`, strings.Join(placeholders, ","))

	var perms []string
	if err := r.db.SelectContext(ctx, &perms, query, args...); err != nil {
		return nil, err
	}
	return perms, nil
}

func (r *AuthRepository) GetCurrentAcademicYearID(schoolID int64) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	var id int64
	if err := r.db.GetContext(ctx, &id, `
		SELECT id FROM academic_years
		WHERE school_id  = ?
		  AND is_current = 1
		LIMIT 1
	`, schoolID); err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *AuthRepository) GetCurrentTermID(academicYearID int64) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	var id int64
	if err := r.db.GetContext(ctx, &id, `
		SELECT id FROM terms
		WHERE academic_year_id = ?
		  AND is_current       = 1
		LIMIT 1
	`, academicYearID); err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *AuthRepository) UpdateLastLogin(userID int64) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()
	//nolint:errcheck
	r.db.ExecContext(ctx,
		`UPDATE users SET last_login_at = NOW() WHERE id = ?`, userID)
}

// BlacklistToken inserts a token JTI into token_blacklist so it cannot be
// reused. Called on logout. The MySQL event in migration 004 auto-purges
// expired rows hourly, keeping the table small.
func (r *AuthRepository) BlacklistToken(jti string, userID int64, expiresAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx,
		`INSERT IGNORE INTO token_blacklist (jti, user_id, expires_at) VALUES (?, ?, ?)`,
		jti, userID, expiresAt)
	return err
}

// IsTokenBlacklisted returns true when the JTI is in the blacklist and has
// not yet expired. Expired rows are harmless but cleaned up by the MySQL event.
func (r *AuthRepository) IsTokenBlacklisted(jti string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()
	var count int
	err := r.db.GetContext(ctx, &count,
		`SELECT COUNT(*) FROM token_blacklist WHERE jti = ? AND expires_at > NOW()`, jti)
	return count > 0, err
}
