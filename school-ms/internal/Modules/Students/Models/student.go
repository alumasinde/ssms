package models

import "time"

type Student struct {
	ID          int64     `db:"id" json:"id"`
	SchoolID    int64     `db:"school_id" json:"school_id"`
	ClassID     int64     `db:"class_id" json:"class_id"`
	AdmissionNo string    `db:"admission_no" json:"admission_no"`

	FirstName   string    `db:"first_name" json:"first_name"`
	MiddleName  string    `db:"middle_name" json:"middle_name"`
	LastName    string    `db:"last_name" json:"last_name"`

	Gender      string    `db:"gender" json:"gender"`
	DOB         string    `db:"dob" json:"dob"`
	PhotoURL    string    `db:"photo_url" json:"photo_url"`

	IsActive    bool      `db:"is_active" json:"is_active"`
	EnrolledAt  time.Time `db:"enrolled_at" json:"enrolled_at"`
}