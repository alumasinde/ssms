package dtos

type CreateTermDTO struct {
	AcademicYearID int64  `json:"academic_year_id"`
	Name           string `json:"name"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	IsCurrent      bool   `json:"is_current"`
}

type UpdateTermDTO struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
