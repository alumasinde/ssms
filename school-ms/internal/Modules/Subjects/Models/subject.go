package models

type Subject struct {
	ID       int64  `db:"id"`
	SchoolID int64  `db:"school_id"`
	Name     string `db:"name"`
	Code     string `db:"code"`
	IsActive bool   `db:"is_active"`
}
