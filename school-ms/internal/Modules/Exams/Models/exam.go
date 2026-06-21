package models

type Exam struct {
	ID        int64  `db:"id"         json:"id"`
	SchoolID  int64  `db:"school_id"  json:"school_id"`
	TermID    int64  `db:"term_id"    json:"term_id"`
	ClassID   *int64 `db:"class_id"   json:"class_id"`
	Name      string `db:"name"       json:"name"`
	Type      string `db:"type"       json:"type"`
	StartDate string `db:"start_date" json:"start_date"`
	EndDate   string `db:"end_date"   json:"end_date"`
}

type ExamResult struct {
	ID        int64   `db:"id"         json:"id"`
	ExamID    int64   `db:"exam_id"    json:"exam_id"`
	StudentID int64   `db:"student_id" json:"student_id"`
	SubjectID int64   `db:"subject_id" json:"subject_id"`
	ClassID   int64   `db:"class_id"   json:"class_id"`
	GradedBy  int64   `db:"graded_by"  json:"graded_by"`
	Marks     float64 `db:"marks"      json:"marks"`
	MaxMarks  float64 `db:"max_marks"  json:"max_marks"`
	Grade     string  `db:"grade"      json:"grade"`
	Remarks   string  `db:"remarks"    json:"remarks"`
}

type GradeScale struct {
	ID       int64   `db:"id"        json:"id"`
	SchoolID int64   `db:"school_id" json:"school_id"`
	Grade    string  `db:"grade"     json:"grade"`
	MinScore float64 `db:"min_score" json:"min_score"`
	MaxScore float64 `db:"max_score" json:"max_score"`
	Remark   string  `db:"remark"    json:"remark"`
}

type ExamResultDetail struct {
	ExamResult
	StudentName string `db:"student_name" json:"student_name"`
	AdmissionNo string `db:"admission_no" json:"admission_no"`
	SubjectName string `db:"subject_name" json:"subject_name"`
	SubjectCode string `db:"subject_code" json:"subject_code"`
}

type StudentReportCard struct {
	StudentID   int64          `json:"student_id"`
	StudentName string         `json:"student_name"`
	AdmissionNo string         `json:"admission_no"`
	ClassName   string         `json:"class_name"`
	Results     []SubjectScore `json:"results"`
	TotalMarks  float64        `json:"total_marks"`
	Average     float64        `json:"average"`
	Position    int            `json:"position"`
}

type SubjectScore struct {
	SubjectName string  `json:"subject_name"`
	Marks       float64 `json:"marks"`
	MaxMarks    float64 `json:"max_marks"`
	Grade       string  `json:"grade"`
	Remarks     string  `json:"remarks"`
}
