package dtos

type CreateTeacherDTO struct {
	UserID        int64  `json:"user_id"`
	SchoolID      int64  `json:"school_id"`
	EmployeeNo    string `json:"employee_no"`
	Phone         string `json:"phone"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	Qualification string `json:"qualification"`
	PhotoURL      string `json:"photo_url"`
}

type AssignSubjectDTO struct {
	TeacherID int64 `json:"teacher_id"`
	SubjectID int64 `json:"subject_id"`
	ClassID   int64 `json:"class_id"`
}
