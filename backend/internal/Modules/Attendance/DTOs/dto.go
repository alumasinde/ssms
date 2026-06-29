package dtos

type MarkAttendanceDTO struct {
	ClassID int64              `json:"class_id"`
	TermID  int64              `json:"term_id"`
	Date    string             `json:"date"`
	Records []AttendanceRecord `json:"records"`
}

type AttendanceRecord struct {
	StudentID int64  `json:"student_id"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
}
