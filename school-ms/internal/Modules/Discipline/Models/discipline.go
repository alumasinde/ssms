package models
import "time"

type DisciplineRecord struct {
	ID           int64     `db:"id"            json:"id"`
	SchoolID     int64     `db:"school_id"     json:"school_id"`
	StudentID    int64     `db:"student_id"    json:"student_id"`
	TermID       int64     `db:"term_id"       json:"term_id"`
	IncidentDate string    `db:"incident_date" json:"incident_date"`
	Type         string    `db:"type"          json:"type"`
	Description  string    `db:"description"   json:"description"`
	ActionTaken  string    `db:"action_taken"  json:"action_taken"`
	RecordedBy   int64     `db:"recorded_by"   json:"recorded_by"`
	CreatedAt    time.Time `db:"created_at"    json:"created_at"`
}

type DisciplineDetail struct {
	DisciplineRecord
	StudentName string `db:"student_name" json:"student_name"`
	AdmissionNo string `db:"admission_no" json:"admission_no"`
	TermName    string `db:"term_name"    json:"term_name"`
	RecorderName string `db:"recorder_name" json:"recorder_name"`
}
