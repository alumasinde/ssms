package dtos

type MarkStaffDTO struct {
	Date    string        `json:"date"`
	Records []StaffEntry `json:"records"`
}
type StaffEntry struct {
	TeacherID int64  `json:"teacher_id"`
	Status    string `json:"status"`
	CheckIn   string `json:"check_in"`
	CheckOut  string `json:"check_out"`
	Remark    string `json:"remark"`
}
