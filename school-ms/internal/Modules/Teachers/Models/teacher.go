package models

type Teacher struct {
	ID            int64  `db:"id"`
	UserID        int64  `db:"user_id"`
	SchoolID      int64  `db:"school_id"`
	EmployeeNo    string `db:"employee_no"`
	Phone         string `db:"phone"`
	Gender        string `db:"gender"`
	DOB           string `db:"dob"`
	Qualification string `db:"qualification"`
	PhotoURL      string `db:"photo_url"`
}

type TeacherSubject struct {
	ID        int64 `db:"id"`
	TeacherID int64 `db:"teacher_id"`
	SubjectID int64 `db:"subject_id"`
	ClassID   int64 `db:"class_id"`
}

type TeacherDetail struct {
	Teacher
	Name  string `db:"name"`
	Email string `db:"email"`
}
