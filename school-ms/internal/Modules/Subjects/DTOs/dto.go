package dtos

type CreateSubjectDTO struct {
	SchoolID int64  `json:"school_id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
}
