package dtos

type CreateClassDTO struct {
	SchoolID int64  `json:"school_id"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Stream   string `json:"stream"`
}
