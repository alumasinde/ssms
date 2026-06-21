package models
import "time"

type Assignment struct {
	ID          int64     `db:"id"          json:"id"`
	SchoolID    int64     `db:"school_id"   json:"school_id"`
	ClassID     int64     `db:"class_id"    json:"class_id"`
	SubjectID   int64     `db:"subject_id"  json:"subject_id"`
	TeacherID   int64     `db:"teacher_id"  json:"teacher_id"`
	TermID      int64     `db:"term_id"     json:"term_id"`
	Title       string    `db:"title"       json:"title"`
	Description string    `db:"description" json:"description"`
	DueDate     string    `db:"due_date"    json:"due_date"`
	MaxMarks    float64   `db:"max_marks"   json:"max_marks"`
	CreatedAt   time.Time `db:"created_at"  json:"created_at"`
}

type AssignmentDetail struct {
	Assignment
	ClassName   string `db:"class_name"   json:"class_name"`
	SubjectName string `db:"subject_name" json:"subject_name"`
	TeacherName string `db:"teacher_name" json:"teacher_name"`
	TermName    string `db:"term_name"    json:"term_name"`
}
