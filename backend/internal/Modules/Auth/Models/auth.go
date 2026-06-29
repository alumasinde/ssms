package models

// LoginUser is the minimal user projection used during authentication.
//
// Schema facts (users table in latest backup):
//   - first_name + last_name columns (no single `name` column)
//   - NO `role` enum column — roles live in user_roles → roles tables
//   - school_id is nullable (superadmin has no school)
//   - soft-delete via deleted_at
type LoginUser struct {
	ID           int64   `db:"id"`
	TenantID     int64   `db:"tenant_id"`
	SchoolID     *int64  `db:"school_id"`
	FirstName    string  `db:"first_name"`
	LastName     string  `db:"last_name"`
	Email        string  `db:"email"`
	PasswordHash string  `db:"password_hash"`
	IsActive     bool    `db:"is_active"`
}

// FullName returns a display-ready concatenated name.
func (u *LoginUser) FullName() string {
	if u.FirstName == "" {
		return u.LastName
	}
	if u.LastName == "" {
		return u.FirstName
	}
	return u.FirstName + " " + u.LastName
}
