package models

type StaffAttendance struct {
	ID         int64  `db:"id"          json:"id"`
	TeacherID  int64  `db:"teacher_id"  json:"teacher_id"`
	SchoolID   int64  `db:"school_id"   json:"school_id"`
	Date       string `db:"date"        json:"date"`
	Status     string `db:"status"      json:"status"`
	CheckIn    string `db:"check_in"    json:"check_in"`
	CheckOut   string `db:"check_out"   json:"check_out"`
	RecordedBy int64  `db:"recorded_by" json:"recorded_by"`
	Remark     string `db:"remark"      json:"remark"`
}

type StaffAttendanceDetail struct {
	StaffAttendance
	TeacherName string `db:"teacher_name" json:"teacher_name"`
	EmployeeNo  string `db:"employee_no"  json:"employee_no"`
}
