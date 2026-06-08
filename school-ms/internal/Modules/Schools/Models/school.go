package models

import "time"

type School struct {
	ID        int64     `db:"id"`
	TenantID  int64     `db:"tenant_id"`
	Name      string    `db:"name"`
	Code      string    `db:"code"`
	Address   string    `db:"address"`
	Phone     string    `db:"phone"`
	Email     string    `db:"email"`
	LogoURL   string    `db:"logo_url"`
	CreatedAt time.Time `db:"created_at"`
}

type AcademicYear struct {
	ID        int64  `db:"id"`
	SchoolID  int64  `db:"school_id"`
	Name      string `db:"name"`
	StartDate string `db:"start_date"`
	EndDate   string `db:"end_date"`
	IsCurrent bool   `db:"is_current"`
}

type Term struct {
	ID             int64  `db:"id"`
	AcademicYearID int64  `db:"academic_year_id"`
	Name           string `db:"name"`
	StartDate      string `db:"start_date"`
	EndDate        string `db:"end_date"`
	IsCurrent      bool   `db:"is_current"`
}
