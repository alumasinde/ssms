package dtos

type CreateSlotDTO struct {
	ClassID   int64  `json:"class_id"`
	SubjectID int64  `json:"subject_id"`
	TeacherID int64  `json:"teacher_id"`
	TermID    int64  `json:"term_id"`
	DayOfWeek int    `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Room      string `json:"room"`
}
