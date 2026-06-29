package models

import "time"

type Parent struct {
	ID         int64  `db:"id" json:"id"`
	UserID     int64  `db:"user_id" json:"user_id"`
	SchoolID   int64  `db:"school_id" json:"school_id"`
	Phone      string `db:"phone" json:"phone"`
	Occupation string `db:"occupation" json:"occupation"`
	Address    string `db:"address" json:"address"`
}

type ParentStudent struct {
	ID           int64  `db:"id" json:"id"`
	ParentID     int64  `db:"parent_id" json:"parent_id"`
	StudentID    int64  `db:"student_id" json:"student_id"`
	Relationship string `db:"relationship" json:"relationship"`
}
type ParentDetail struct {
	Parent

	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name"  json:"last_name"`
	Email     string `db:"email"      json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}