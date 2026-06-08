package models

type AcademicYear struct {
	ID       int64  `db:"id"`
	SchoolID int64  `db:"school_id"`
	Name     string `db:"name"`
	StartDate string `db:"start_date"`
	EndDate   string `db:"end_date"`
	IsCurrent bool   `db:"is_current"`
}
