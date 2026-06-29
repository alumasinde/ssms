package repositories
import (
	"context"
	"strings"
	"time"

	models "school-ms/internal/Modules/Users/Models"

	"github.com/jmoiron/sqlx"
)

const dbTimeout = 5 * time.Second

// userCols matches models.User exactly — no deleted_at/deleted_by in struct.
const userCols = `u.id, u.tenant_id, u.school_id,
	u.first_name, u.last_name, u.email, u.password_hash,
	u.phone, u.avatar_url, u.is_active,
	u.last_login_at, u.created_at, u.updated_at`

type UserRepository struct{ db *sqlx.DB }

func NewUserRepository(db *sqlx.DB) *UserRepository { return &UserRepository{db: db} }

func (r *UserRepository) Create(u *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	res, err := r.db.ExecContext(ctx, `
		INSERT INTO users (tenant_id,school_id,first_name,last_name,email,password_hash,phone,is_active)
		VALUES (?,?,?,?,?,?,?,1)
	`, u.TenantID, u.SchoolID, u.FirstName, u.LastName, u.Email, u.PasswordHash, u.Phone)
	if err != nil { return err }
	id, _ := res.LastInsertId(); u.ID = id; return nil
}

// AssignRole resolves roleCode → roles.id (scoped to tenant, falling back to global) then inserts user_roles.
func (r *UserRepository) AssignRole(userID, tenantID int64, roleCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var roleID int64
	if err := r.db.GetContext(ctx, &roleID, `
		SELECT id FROM roles
		WHERE code = ?
		  AND (tenant_id = ? OR tenant_id IS NULL)
		  AND is_active = 1
		ORDER BY tenant_id DESC LIMIT 1
	`, roleCode, tenantID); err != nil {
		return err
	}
	_, err := r.db.ExecContext(ctx,
		`INSERT IGNORE INTO user_roles (user_id,role_id) VALUES (?,?)`, userID, roleID)
	return err
}

// ReplaceRole removes all existing roles and assigns a new one in a transaction.
func (r *UserRepository) ReplaceRole(userID, tenantID int64, roleCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var roleID int64
	if err := r.db.GetContext(ctx, &roleID, `
		SELECT id FROM roles
		WHERE code = ?
		  AND (tenant_id = ? OR tenant_id IS NULL)
		  AND is_active = 1
		ORDER BY tenant_id DESC LIMIT 1
	`, roleCode, tenantID); err != nil {
		return err
	}
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil { return err }
	if _, err := tx.Exec(`DELETE FROM user_roles WHERE user_id=?`, userID); err != nil {
		tx.Rollback(); return err
	}
	if _, err := tx.Exec(`INSERT INTO user_roles (user_id,role_id) VALUES (?,?)`, userID, roleID); err != nil {
		tx.Rollback(); return err
	}
	return tx.Commit()
}

// FindRoleNames returns roles.name[] for a user via user_roles join.
func (r *UserRepository) FindRoleNames(userID int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var roles []string
	err := r.db.SelectContext(ctx, &roles, `
		SELECT ro.name FROM roles ro
		INNER JOIN user_roles ur ON ur.role_id = ro.id
		WHERE ur.user_id = ? AND ro.is_active = 1
		ORDER BY ro.name
	`, userID)
	return roles, err
}

func (r *UserRepository) FindByEmail(email string, tenantID int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var u models.User
	err := r.db.GetContext(ctx, &u, `
		SELECT `+userCols+` FROM users u
		WHERE u.email=? AND u.tenant_id=? AND u.is_active=1
		  AND (u.deleted_at IS NULL OR u.deleted_at > NOW()) LIMIT 1
	`, strings.TrimSpace(strings.ToLower(email)), tenantID)
	return &u, err
}

func (r *UserRepository) FindByID(id int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var u models.User
	err := r.db.GetContext(ctx, &u, `
		SELECT `+userCols+` FROM users u
		WHERE u.id=? AND (u.deleted_at IS NULL OR u.deleted_at > NOW())
	`, id)
	return &u, err
}

func (r *UserRepository) ListByTenant(tenantID int64) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var list []models.User
	err := r.db.SelectContext(ctx, &list, `
		SELECT `+userCols+` FROM users u
		WHERE u.tenant_id=? AND u.is_active=1
		  AND (u.deleted_at IS NULL OR u.deleted_at > NOW())
		ORDER BY u.first_name, u.last_name
	`, tenantID)
	return list, err
}

func (r *UserRepository) ListBySchool(schoolID int64) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var list []models.User
	err := r.db.SelectContext(ctx, &list, `
		SELECT `+userCols+` FROM users u
		WHERE u.school_id=? AND u.is_active=1
		  AND (u.deleted_at IS NULL OR u.deleted_at > NOW())
		ORDER BY u.first_name, u.last_name
	`, schoolID)
	return list, err
}

func (r *UserRepository) UpdatePassword(id int64, hash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `UPDATE users SET password_hash=? WHERE id=?`, hash, id)
	return err
}

// SoftDelete stamps deleted_at and sets is_active=0; records who deleted.
func (r *UserRepository) SoftDelete(id, deletedBy int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		UPDATE users SET is_active=0, deleted_at=NOW(), deleted_by=? WHERE id=?
	`, deletedBy, id)
	return err
}

func (r *UserRepository) Activate(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx,
		`UPDATE users SET is_active=1, deleted_at=NULL, deleted_by=NULL WHERE id=?`, id)
	return err
}