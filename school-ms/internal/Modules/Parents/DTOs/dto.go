package dtos

type CreateParentDTO struct {
	UserID     int64  `json:"user_id"`
	SchoolID   int64  `json:"school_id"`
	Phone      string `json:"phone"`
	Occupation string `json:"occupation"`
	Address    string `json:"address"`
}

type LinkStudentDTO struct {
	ParentID     int64  `json:"parent_id"`
	StudentID    int64  `json:"student_id"`
	Relationship string `json:"relationship"`
}
