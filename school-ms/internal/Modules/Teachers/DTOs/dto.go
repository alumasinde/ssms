package dtos
import "strings"

// CreateTeacherDTO — POST /teachers.
type CreateTeacherDTO struct {
UserID         int64   `json:"user_id"`
SchoolID       int64   `json:"school_id"`   // injected from context
EmployeeNo     string  `json:"employee_no"`
Phone          *string `json:"phone"`
Gender         *string `json:"gender"`
DOB            *string `json:"dob"`
TSCNo          *string `json:"tsc_no"`
HireDate       *string `json:"hire_date"`
Qualification  *string `json:"qualification"`
Specialization *string `json:"specialization"`
EmploymentType string  `json:"employment_type"` // permanent|contract|part_time
NationalID     *string `json:"national_id"`
Address        *string `json:"address"`
PhotoURL       *string `json:"photo_url"`
}

func (d *CreateTeacherDTO) Validate() map[string]string {
errs := map[string]string{}
if d.UserID == 0 { errs["user_id"] = "user_id is required" }
if strings.TrimSpace(d.EmployeeNo) == "" { errs["employee_no"] = "employee number is required" }
valid := map[string]bool{"permanent": true, "contract": true, "part_time": true, "": true}
if !valid[d.EmploymentType] { errs["employment_type"] = "must be permanent, contract, or part_time" }
if len(errs) > 0 { return errs }
return nil
}

// UpdateTeacherDTO — PUT /teachers/{id}. All fields are optional (patch semantics).
type UpdateTeacherDTO struct {
Phone          *string `json:"phone"`
Gender         *string `json:"gender"`
DOB            *string `json:"dob"`
TSCNo          *string `json:"tsc_no"`
HireDate       *string `json:"hire_date"`
Qualification  *string `json:"qualification"`
Specialization *string `json:"specialization"`
NationalID     *string `json:"national_id"`
Address        *string `json:"address"`
PhotoURL       *string `json:"photo_url"`
IsClassTeacher *bool   `json:"is_class_teacher"`
EmploymentType *string `json:"employment_type"`
}

// AssignSubjectDTO — POST /teachers/{id}/subjects.
// teacher_id is read from the URL param, NOT the body.
type AssignSubjectDTO struct {
SubjectID int64 `json:"subject_id"`
ClassID   int64 `json:"class_id"`
}

func (d *AssignSubjectDTO) Validate() map[string]string {
errs := map[string]string{}
if d.SubjectID == 0 { errs["subject_id"] = "subject_id is required" }
if d.ClassID == 0   { errs["class_id"]   = "class_id is required" }
if len(errs) > 0 { return errs }
return nil
}