package dtos

type CreateAssignmentDTO struct {
	ClassID     int64   `json:"class_id"`
	SubjectID   int64   `json:"subject_id"`
	TeacherID   int64   `json:"teacher_id"`
	TermID      int64   `json:"term_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	DueDate     string  `json:"due_date"`
	MaxMarks    float64 `json:"max_marks"`
}
