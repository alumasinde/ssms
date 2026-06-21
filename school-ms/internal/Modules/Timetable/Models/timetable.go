package models

type TimetableSlot struct {
	ID        int64  `db:"id"          json:"id"`
	SchoolID  int64  `db:"school_id"   json:"school_id"`
	ClassID   int64  `db:"class_id"    json:"class_id"`
	SubjectID int64  `db:"subject_id"  json:"subject_id"`
	TeacherID int64  `db:"teacher_id"  json:"teacher_id"`
	TermID    int64  `db:"term_id"     json:"term_id"`
	DayOfWeek int    `db:"day_of_week" json:"day_of_week"`
	StartTime string `db:"start_time"  json:"start_time"`
	EndTime   string `db:"end_time"    json:"end_time"`
	Room      string `db:"room"        json:"room"`
}

type TimetableSlotDetail struct {
	TimetableSlot
	ClassName   string `db:"class_name"   json:"class_name"`
	SubjectName string `db:"subject_name" json:"subject_name"`
	SubjectCode string `db:"subject_code" json:"subject_code"`
	TeacherName string `db:"teacher_name" json:"teacher_name"`
}
