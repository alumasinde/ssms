package models

import "time"

type Subject struct {
	ID       int64  `db:"id" json:"id"`
	SchoolID int64  `db:"school_id" json:"school_id"`
	Name     string `db:"name" json:"name"`
	Code     string `db:"code" json:"code"`
	IsActive bool   `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
