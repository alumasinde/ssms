package models

import "time"

// Student mirrors the live students table (all columns).
type Student struct {
ID          int64      `db:"id"           json:"id"`
SchoolID    int64      `db:"school_id"    json:"school_id"`
ClassID     int64      `db:"class_id"     json:"class_id"`
AdmissionNo string     `db:"admission_no" json:"admission_no"`
FirstName   string     `db:"first_name"   json:"first_name"`
MiddleName  *string    `db:"middle_name"  json:"middle_name,omitempty"`
LastName    string     `db:"last_name"    json:"last_name"`
Gender      *string    `db:"gender"       json:"gender,omitempty"`
DOB         *time.Time `db:"dob"          json:"dob,omitempty"`
Nationality *string    `db:"nationality"  json:"nationality,omitempty"`
NationalID  *string    `db:"national_id"  json:"national_id,omitempty"`
Religion    *string    `db:"religion"     json:"religion,omitempty"`
BloodGroup  *string    `db:"blood_group"  json:"blood_group,omitempty"`
Address     *string    `db:"address"      json:"address,omitempty"`
MedicalNotes *string   `db:"medical_notes" json:"medical_notes,omitempty"`
PhotoURL    string     `db:"photo_url"    json:"photo_url"`
IsActive    bool       `db:"is_active"    json:"is_active"`
EnrolledAt  time.Time  `db:"enrolled_at"  json:"enrolled_at"`
LeftDate    *time.Time `db:"left_date"    json:"left_date,omitempty"`
LeftReason  *string    `db:"left_reason"  json:"left_reason,omitempty"`
}

func (s *Student) FullName() string {
if s.MiddleName != nil && *s.MiddleName != "" {
return s.FirstName + " " + *s.MiddleName + " " + s.LastName
}
return s.FirstName + " " + s.LastName
}