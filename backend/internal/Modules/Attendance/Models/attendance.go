package models

type Attendance struct {
	ID         int64  `db:"id"`
	StudentID  int64  `db:"student_id"`
	ClassID    int64  `db:"class_id"`
	TermID     int64  `db:"term_id"`
	RecordedBy int64  `db:"recorded_by"`
	Date       string `db:"date"`
	Status     string `db:"status"` // present, absent, late, excused
	Remark     string `db:"remark"`
}

type AttendanceSummary struct {
	StudentID   int64   `db:"student_id"`
	StudentName string  `db:"student_name"`
	Total       int     `db:"total"`
	Present     int     `db:"present"`
	Absent      int     `db:"absent"`
	Late        int     `db:"late"`
	Percent     float64 `db:"percent"`
}
