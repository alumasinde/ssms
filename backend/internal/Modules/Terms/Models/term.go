package models

type Term struct {
	ID             int64  `db:"id"               json:"id"`
	AcademicYearID int64  `db:"academic_year_id" json:"academic_year_id"`
	SchoolID       int64  `db:"school_id"        json:"school_id"`
	Name           string `db:"name"             json:"name"`
	StartDate      string `db:"start_date"       json:"start_date"`
	EndDate        string `db:"end_date"         json:"end_date"`
	IsCurrent      bool   `db:"is_current"       json:"is_current"`
}
