package models

type Subject struct {
	ID       int64  `db:"id" json:"id"`
	SchoolID int64  `db:"school_id" json:"school_id"`
	Name     string `db:"name" json:"name"`
	Code     string `db:"code" json:"code"`
	IsActive bool   `db:"is_active" json:"is_active"`
}
