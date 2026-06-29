package repositories

import (
	"school-ms/internal/Modules/Tenants/Models"

	"github.com/jmoiron/sqlx"
)

type TenantRepository struct {
	db *sqlx.DB
}

func NewTenantRepository(db *sqlx.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) Create(t *models.Tenant) error {
	q := `INSERT INTO tenants (slug, name, domain, plan, is_active) VALUES (?,?,?,?,1)`
	res, err := r.db.Exec(q, t.Slug, t.Name, t.Domain, t.Plan)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	t.ID = id
	return nil
}

func (r *TenantRepository) FindByID(id int64) (*models.Tenant, error) {
	var t models.Tenant
	err := r.db.Get(&t, `SELECT * FROM tenants WHERE id = ? AND is_active=1`, id)
	return &t, err
}

func (r *TenantRepository) FindBySlug(slug string) (*models.Tenant, error) {
	var t models.Tenant
	err := r.db.Get(&t, `SELECT * FROM tenants WHERE slug = ? AND is_active=1`, slug)
	return &t, err
}

// FindByDomain resolves a tenant from a full domain or subdomain.
// It tries an exact match on the domain column first, then falls back
// to matching the subdomain (e.g. "highwayhigh" from "ssms.highwayhigh.ac.ke").
func (r *TenantRepository) FindByDomain(host string) (*models.Tenant, error) {
	var t models.Tenant
	// Exact domain match (e.g. highwayhighschool.ac.ke stored in domain column)
	err := r.db.Get(&t,
		`SELECT * FROM tenants WHERE domain = ? AND is_active=1 LIMIT 1`, host)
	if err == nil {
		return &t, nil
	}
	// Slug match from subdomain: ssms.highwayhigh.ac.ke → slug 'highwayhigh'
	err = r.db.Get(&t,
		`SELECT * FROM tenants WHERE slug = ? AND is_active=1 LIMIT 1`, host)
	return &t, err
}

func (r *TenantRepository) List() ([]models.Tenant, error) {
	var list []models.Tenant
	err := r.db.Select(&list, `SELECT * FROM tenants ORDER BY created_at DESC`)
	return list, err
}

func (r *TenantRepository) Update(id int64, name, domain, plan string, isActive bool) error {
	_, err := r.db.Exec(
		`UPDATE tenants SET name=?, domain=?, plan=?, is_active=? WHERE id=?`,
		name, domain, plan, isActive, id,
	)
	return err
}

func (r *TenantRepository) Delete(id int64) error {
	_, err := r.db.Exec(`UPDATE tenants SET is_active=0 WHERE id=?`, id)
	return err
}
