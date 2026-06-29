package dtos

import "strings"

// CreateClassDTO is used for POST /classes.
type CreateClassDTO struct {
	SchoolID int64   `json:"school_id"` // injected from context, never from client
	Name     string  `json:"name"`
	Level    string  `json:"level"`
	Stream   *string `json:"stream"`    // nullable
}

// Validate returns a map of field → message when the DTO is invalid, nil otherwise.
func (d *CreateClassDTO) Validate() map[string]string {
	errs := map[string]string{}
	if strings.TrimSpace(d.Name) == "" {
		errs["name"] = "class name is required"
	}
	if strings.TrimSpace(d.Level) == "" {
		errs["level"] = "class level is required"
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// UpdateClassDTO is used for PUT /classes/{id}.
// SchoolID is not updated — only name, level, stream.
type UpdateClassDTO struct {
	Name   string  `json:"name"`
	Level  string  `json:"level"`
	Stream *string `json:"stream"`
}

func (d *UpdateClassDTO) Validate() map[string]string {
	errs := map[string]string{}
	if strings.TrimSpace(d.Name) == "" {
		errs["name"] = "class name is required"
	}
	if strings.TrimSpace(d.Level) == "" {
		errs["level"] = "class level is required"
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// AssignSubjectDTO is used for POST /classes/{id}/subjects.
type AssignSubjectDTO struct {
	SubjectID    int64 `json:"subject_id"`
	IsCompulsory bool  `json:"is_compulsory"`
}
