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

// FindUserByEmailAndTenant returns the minimal login projection.
//
// Schema notes:
//   - users has first_name + last_name (no composite `name` column)
//   - users has no `role` column — roles are in user_roles → roles
//   - soft-delete guarded by deleted_at IS NULL OR deleted_at > NOW()
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

// FindRoleNamesByUserID returns the role names (roles.name) assigned to a user.
//
// Schema: user_roles(user_id, role_id) → roles(id, name, is_active)
// We return roles.name because that is what Claims.Roles []string stores and
// what middleware.GetRoles() exposes to handlers.
func (r *AuthRepository) FindRoleNamesByUserID(userID int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	var roles []string
	err := r.db.SelectContext(ctx, &roles, `
		SELECT r.name
		FROM roles r
		INNER JOIN user_roles ur ON ur.role_id = r.id
		WHERE ur.user_id  = ?
		  AND r.is_active = 1
		ORDER BY r.name
	`, userID)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// FindPermissionsByRoleNames returns all distinct permission names for a slice
// of role names.
//
// Schema: role_permissions(role_id, permission_id) → permissions(id, name)
//         roles(id, name) — joined via role_id FK
//
// Uses a dynamic IN(...) clause built from the role names slice.
// Returns an empty slice (not an error) when roles is empty.
func (r *AuthRepository) FindPermissionsByRoleNames(roleNames []string) ([]string, error) {
	if len(roleNames) == 0 {
		return []string{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	// Build  WHERE r.name IN (?, ?, ...)  safely
	placeholders := make([]string, len(roleNames))
	args := make([]interface{}, len(roleNames))
	for i, name := range roleNames {
		placeholders[i] = "?"
		args[i] = name
	}

	query := fmt.Sprintf(`
		SELECT DISTINCT p.name
		FROM permissions p
		INNER JOIN role_permissions rp ON rp.permission_id = p.id
		INNER JOIN roles r             ON r.id             = rp.role_id
		WHERE r.name      IN (%s)
		  AND r.is_active = 1
		ORDER BY p.name
	`, strings.Join(placeholders, ","))

	var perms []string
	if err := r.db.SelectContext(ctx, &perms, query, args...); err != nil {
		return nil, err
	}
	return perms, nil
}

// GetCurrentAcademicYearID returns the active academic year ID for a school.
// The DB enforces at most one current year via a generated unique column.
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

// GetCurrentTermID returns the active term ID for an academic year.
// The DB enforces at most one current term via a generated unique column.
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

// UpdateLastLogin stamps last_login_at. Fire-and-forget — call in a goroutine.
func (r *AuthRepository) UpdateLastLogin(userID int64) {
	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()
	//nolint:errcheck
	r.db.ExecContext(ctx,
		`UPDATE users SET last_login_at = NOW() WHERE id = ?`, userID)
}
