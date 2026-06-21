package repositories

import (
	models "school-ms/internal/Modules/StaffAttendance/Models"
	"github.com/jmoiron/sqlx"
)

type StaffAttendanceRepository struct{ db *sqlx.DB }
func NewStaffAttendanceRepository(db *sqlx.DB) *StaffAttendanceRepository { return &StaffAttendanceRepository{db:db} }

func (r *StaffAttendanceRepository) BulkUpsert(records []models.StaffAttendance) error {
	tx,err := r.db.Beginx()
	if err != nil { return err }
	q := `INSERT INTO staff_attendance (teacher_id,school_id,date,status,check_in,check_out,recorded_by,remark)
	      VALUES (?,?,?,?,?,?,?,?)
	      ON DUPLICATE KEY UPDATE status=VALUES(status),check_in=VALUES(check_in),
	      check_out=VALUES(check_out),remark=VALUES(remark),recorded_by=VALUES(recorded_by)`
	for _,a := range records {
		if _,err := tx.Exec(q,a.TeacherID,a.SchoolID,a.Date,a.Status,a.CheckIn,a.CheckOut,a.RecordedBy,a.Remark); err != nil {
			tx.Rollback(); return err
		}
	}
	return tx.Commit()
}
func (r *StaffAttendanceRepository) ListByDate(schoolID int64, date string) ([]models.StaffAttendanceDetail,error) {
	var list []models.StaffAttendanceDetail
	return list, r.db.Select(&list,
		`SELECT sa.*,u.name AS teacher_name,t.employee_no
		 FROM staff_attendance sa
		 JOIN teachers t ON t.id=sa.teacher_id
		 JOIN users    u ON u.id=t.user_id
		 WHERE sa.school_id=? AND sa.date=? ORDER BY u.name`,schoolID,date)
}
func (r *StaffAttendanceRepository) SummaryByMonth(schoolID int64, yearMonth string) ([]models.StaffAttendanceDetail,error) {
	var list []models.StaffAttendanceDetail
	return list, r.db.Select(&list,
		`SELECT sa.teacher_id,sa.school_id,sa.status,
		        '' AS date,'' AS check_in,'' AS check_out,0 AS recorded_by,'' AS remark,0 AS id,
		        u.name AS teacher_name, t.employee_no,
		        COUNT(*) AS total_days,
		        SUM(sa.status='present') AS present_days
		 FROM staff_attendance sa
		 JOIN teachers t ON t.id=sa.teacher_id
		 JOIN users    u ON u.id=t.user_id
		 WHERE sa.school_id=? AND DATE_FORMAT(sa.date,'%Y-%m')=?
		 GROUP BY sa.teacher_id,u.name,t.employee_no`,schoolID,yearMonth)
}
