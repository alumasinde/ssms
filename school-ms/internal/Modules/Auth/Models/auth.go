package models

// LoginUser is the minimal user projection used during login.
// Full user data lives in the Users module.
type LoginUser struct {
	ID           int64  `db:"id"`
	TenantID     int64  `db:"tenant_id"`
	Name         string `db:"name"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	Role         string `db:"role"`
	SchoolID     int64  `db:"school_id"` // first school for this user's tenant
	IsActive     bool   `db:"is_active"`
}

type Permission struct {
    ID   int64  `db:"id"   json:"id"`
    Name string `db:"name" json:"name"`
}