package repositories

import (
	"school-ms/internal/Modules/Auth/Models"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// FindUserByEmailAndTenant locates an active user by email within a resolved tenant.
// It also fetches the first school belonging to the tenant so school_id can be
// embedded in the JWT without requiring a second round-trip.
func (r *AuthRepository) FindUserByEmailAndTenant(email string, tenantID int64) (*models.LoginUser, error) {
	var user models.LoginUser
	query := `
		SELECT
			u.id,
			u.tenant_id,
			u.name,
			u.email,
			u.password_hash,
			u.role,
			u.is_active,
			COALESCE(
				(SELECT id FROM schools WHERE tenant_id = u.tenant_id ORDER BY id LIMIT 1),
				0
			) AS school_id
		FROM users u
		WHERE u.email      = ?
		  AND u.tenant_id  = ?
		  AND u.is_active  = 1
		LIMIT 1
	`
	err := r.db.Get(&user, query, email, tenantID)
	return &user, err
}

func (r *AuthRepository) FindUserPermissions(role string) ([]string, error) {
    var perms []string
    err := r.db.Select(&perms, `
        SELECT p.name 
        FROM permissions p
        JOIN role_permissions rp ON rp.permission_id = p.id
        WHERE rp.role = ?`, role)
    return perms, err
}