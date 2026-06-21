package dtos

type CreateDisciplineDTO struct {
	StudentID    int64  `json:"student_id"`
	TermID       int64  `json:"term_id"`
	IncidentDate string `json:"incident_date"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	ActionTaken  string `json:"action_taken"`
}
