package models
import "time"

// Teacher mirrors the live teachers table (all columns).
type Teacher struct {
ID             int64      `db:"id"               json:"id"`
UserID         int64      `db:"user_id"          json:"user_id"`
SchoolID       int64      `db:"school_id"        json:"school_id"`
EmployeeNo     string     `db:"employee_no"      json:"employee_no"`
Phone          *string    `db:"phone"            json:"phone,omitempty"`
Gender         *string    `db:"gender"           json:"gender,omitempty"`
DOB            *string    `db:"dob"              json:"dob,omitempty"`
Qualification  *string    `db:"qualification"    json:"qualification,omitempty"`
TSCNo          *string    `db:"tsc_no"           json:"tsc_no,omitempty"`
Specialization *string    `db:"specialization"   json:"specialization,omitempty"`
HireDate       *string    `db:"hire_date"        json:"hire_date,omitempty"`
EmploymentType string     `db:"employment_type"  json:"employment_type"`
IsClassTeacher bool       `db:"is_class_teacher" json:"is_class_teacher"`
NationalID     *string    `db:"national_id"      json:"national_id,omitempty"`
Address        *string    `db:"address"          json:"address,omitempty"`
IsActive       bool       `db:"is_active"        json:"is_active"`
PhotoURL       *string    `db:"photo_url"        json:"photo_url,omitempty"`
CreatedAt      time.Time  `db:"created_at"       json:"created_at"`
UpdatedAt      time.Time  `db:"updated_at"       json:"updated_at"`
}

// TeacherDetail enriches Teacher with identity from users.
// users has first_name + last_name (no `name` column).
type TeacherDetail struct {
Teacher
FirstName string `db:"first_name" json:"first_name"`
LastName  string `db:"last_name"  json:"last_name"`
Email     string `db:"email"      json:"email"`
}

func (t *TeacherDetail) FullName() string {
if t.FirstName == "" { return t.LastName }
if t.LastName == ""  { return t.FirstName }
return t.FirstName + " " + t.LastName
}

// TeacherSubject mirrors teacher_subjects (includes school_id NOT NULL).
type TeacherSubject struct {
ID        int64 `db:"id"         json:"id"`
TeacherID int64 `db:"teacher_id" json:"teacher_id"`
SubjectID int64 `db:"subject_id" json:"subject_id"`
ClassID   int64 `db:"class_id"   json:"class_id"`
SchoolID  int64 `db:"school_id"  json:"school_id"`
}

type TeacherSubjectDetail struct {
TeacherSubject
SubjectName string `db:"subject_name" json:"subject_name"`
SubjectCode string `db:"subject_code" json:"subject_code"`
ClassName   string `db:"class_name"   json:"class_name"`
}