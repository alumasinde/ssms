package dtos

type CreateStudentDTO struct {
	SchoolID    int64  `json:"school_id"`
	ClassID     int64  `json:"class_id"`
	AdmissionNo string `json:"admission_no"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	DOB         string `json:"dob"`
	PhotoURL    string `json:"photo_url"`
}

type TransferClassDTO struct {
	NewClassID int64 `json:"new_class_id"`
}
