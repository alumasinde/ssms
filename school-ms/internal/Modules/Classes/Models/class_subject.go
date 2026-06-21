package models

// ClassSubject mirrors class_subjects joined with subjects.
//
// Schema: class_subjects(id, class_id, subject_id, is_compulsory, school_id)
//   - school_id is NOT NULL with FK to schools — must be supplied on INSERT
type ClassSubject struct {
	ID           int64  `db:"id"            json:"id"`
	ClassID      int64  `db:"class_id"      json:"class_id"`
	SubjectID    int64  `db:"subject_id"    json:"subject_id"`
	SchoolID     int64  `db:"school_id"     json:"school_id"`
	IsCompulsory bool   `db:"is_compulsory" json:"is_compulsory"`
	SubjectName  string `db:"subject_name"  json:"subject_name"`
	SubjectCode  string `db:"subject_code"  json:"subject_code"`
}
