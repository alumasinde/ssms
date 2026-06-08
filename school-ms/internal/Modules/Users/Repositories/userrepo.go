package repositories

import (
	models "school-ms/internal/Modules/Users/Models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *models.User) error {
	q := `INSERT INTO users (tenant_id, name, email, password_hash, role, is_active) VALUES (?,?,?,?,?,1)`
	res, err := r.db.Exec(q, u.TenantID, u.Name, u.Email, u.PasswordHash, u.Role)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	u.ID = id
	return nil
}

func (r *UserRepository) FindByEmail(email string, tenantID int64) (*models.User, error) {
	var u models.User
	err := r.db.Get(&u, `SELECT * FROM users WHERE email=? AND tenant_id=? AND is_active=1`, email, tenantID)
	return &u, err
}

func (r *UserRepository) FindByID(id int64) (*models.User, error) {
	var u models.User
	err := r.db.Get(&u, `SELECT * FROM users WHERE id=?`, id)
	return &u, err
}

func (r *UserRepository) ListByTenant(tenantID int64) ([]models.User, error) {
	var list []models.User
	err := r.db.Select(&list, `SELECT * FROM users WHERE tenant_id=? ORDER BY name`, tenantID)
	return list, err
}

func (r *UserRepository) UpdatePassword(id int64, hash string) error {
	_, err := r.db.Exec(`UPDATE users SET password_hash=? WHERE id=?`, hash, id)
	return err
}

func (r *UserRepository) Deactivate(id int64) error {
	_, err := r.db.Exec(`UPDATE users SET is_active=0 WHERE id=?`, id)
	return err
}
