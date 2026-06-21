package dtos

import (
"strings"
"time"
)

// CreateStudentDTO — POST /students.
type CreateStudentDTO struct {
SchoolID     int64      `json:"school_id"`   // injected from context
ClassID      int64      `json:"class_id"`
AdmissionNo  string     `json:"admission_no"`
FirstName    string     `json:"first_name"`
MiddleName   *string    `json:"middle_name"`
LastName     string     `json:"last_name"`
Gender       *string    `json:"gender"`
DOB          *time.Time `json:"dob"`
Nationality  *string    `json:"nationality"`
NationalID   *string    `json:"national_id"`
Religion     *string    `json:"religion"`
BloodGroup   *string    `json:"blood_group"`
Address      *string    `json:"address"`
MedicalNotes *string    `json:"medical_notes"`
PhotoURL     string     `json:"photo_url"`
}

func (d *CreateStudentDTO) Validate() map[string]string {
errs := map[string]string{}
if strings.TrimSpace(d.AdmissionNo) == "" { errs["admission_no"] = "admission number is required" }
if strings.TrimSpace(d.FirstName) == ""   { errs["first_name"]   = "first name is required" }
if strings.TrimSpace(d.LastName) == ""    { errs["last_name"]    = "last name is required" }
if d.ClassID == 0                         { errs["class_id"]     = "class_id is required" }
if len(errs) > 0 { return errs }
return nil
}

// UpdateStudentDTO — PUT /students/{id}. Uses patch semantics.
type UpdateStudentDTO struct {
ClassID      int64      `json:"class_id"`
FirstName    string     `json:"first_name"`
MiddleName   *string    `json:"middle_name"`
LastName     string     `json:"last_name"`
Gender       *string    `json:"gender"`
DOB          *time.Time `json:"dob"`
Nationality  *string    `json:"nationality"`
NationalID   *string    `json:"national_id"`
Religion     *string    `json:"religion"`
BloodGroup   *string    `json:"blood_group"`
Address      *string    `json:"address"`
MedicalNotes *string    `json:"medical_notes"`
PhotoURL     string     `json:"photo_url"`
}

func (d *UpdateStudentDTO) Validate() map[string]string {
errs := map[string]string{}
if strings.TrimSpace(d.FirstName) == "" { errs["first_name"] = "first name is required" }
if strings.TrimSpace(d.LastName) == ""  { errs["last_name"]  = "last name is required" }
if len(errs) > 0 { return errs }
return nil
}

// PromoteStudentsDTO — POST /students/promote.
type PromoteStudentsDTO struct {
FromClassID     int64   `json:"from_class_id"`
ToClassID       int64   `json:"to_class_id"`
AcademicYearID  int64   `json:"academic_year_id"`  // for student_class_history
StudentIDs      []int64 `json:"student_ids"`         // empty = promote entire class
}

func (d *PromoteStudentsDTO) Validate() map[string]string {
errs := map[string]string{}
if d.FromClassID == 0    { errs["from_class_id"]    = "from_class_id is required" }
if d.ToClassID == 0      { errs["to_class_id"]      = "to_class_id is required" }
if d.AcademicYearID == 0 { errs["academic_year_id"] = "academic_year_id is required" }
if d.FromClassID == d.ToClassID { errs["to_class_id"] = "target class must differ from source class" }
if len(errs) > 0 { return errs }
return nil
}