package models

import "time"

type Tenant struct {
	ID        int64     `db:"id"`
	Slug      string    `db:"slug"`
	Name      string    `db:"name"`
	Domain    string    `db:"domain"`
	Plan      string    `db:"plan"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
}
