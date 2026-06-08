package dtos

type CreateExamDTO struct {
	SchoolID  int64  `json:"school_id"`
	TermID    int64  `json:"term_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type SubmitResultDTO struct {
	ExamID    int64         `json:"exam_id"`
	Results   []ResultEntry `json:"results"`
}

type ResultEntry struct {
	StudentID int64   `json:"student_id"`
	SubjectID int64   `json:"subject_id"`
	Marks     float64 `json:"marks"`
	MaxMarks  float64 `json:"max_marks"`
	Remarks   string  `json:"remarks"`
}

type CreateGradeScaleDTO struct {
	SchoolID int64   `json:"school_id"`
	Grade    string  `json:"grade"`
	MinScore float64 `json:"min_score"`
	MaxScore float64 `json:"max_score"`
	Remark   string  `json:"remark"`
}
