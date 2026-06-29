package models

import "time"

type User struct {
ID           int64      `db:"id"            json:"id"`
TenantID     int64      `db:"tenant_id"     json:"tenant_id"`
SchoolID     *int64     `db:"school_id"     json:"school_id,omitempty"`
FirstName    string     `db:"first_name"    json:"first_name"`
LastName     string     `db:"last_name"     json:"last_name"`
Email        string     `db:"email"         json:"email"`
PasswordHash string     `db:"password_hash" json:"-"`
Phone        *string    `db:"phone"         json:"phone,omitempty"`
AvatarURL    *string    `db:"avatar_url"    json:"avatar_url,omitempty"`
IsActive     bool       `db:"is_active"     json:"is_active"`
LastLoginAt  *time.Time `db:"last_login_at" json:"last_login_at,omitempty"`
CreatedAt    time.Time  `db:"created_at"    json:"created_at"`
UpdatedAt    time.Time  `db:"updated_at"    json:"updated_at"`
}

func (u *User) FullName() string {
if u.FirstName == "" { return u.LastName }
if u.LastName == ""  { return u.FirstName }
return u.FirstName + " " + u.LastName
}

type UserWithRoles struct {
User
Roles []string `json:"roles"`
}