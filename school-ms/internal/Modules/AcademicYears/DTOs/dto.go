package dtos

type CreateAcademicYearDTO struct {
	SchoolID int64  `json:"school_id"`
	Name     string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	IsCurrent bool   `json:"is_current"`
}
