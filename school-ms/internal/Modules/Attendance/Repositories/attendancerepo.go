package repositories

import (
	models "school-ms/internal/Modules/Attendance/Models"
	"github.com/jmoiron/sqlx"
)

type AttendanceRepository struct{ db *sqlx.DB }

func NewAttendanceRepository(db *sqlx.DB) *AttendanceRepository { return &AttendanceRepository{db: db} }

func (r *AttendanceRepository) BulkUpsert(records []models.Attendance) error {
	tx, err := r.db.Beginx()
	if err != nil { return err }
	q := `INSERT INTO attendance (student_id,class_id,term_id,recorded_by,date,status,remark)
	      VALUES (?,?,?,?,?,?,?)
	      ON DUPLICATE KEY UPDATE status=VALUES(status),remark=VALUES(remark),recorded_by=VALUES(recorded_by)`
	for _, a := range records {
		if _, err := tx.Exec(q, a.StudentID, a.ClassID, a.TermID, a.RecordedBy, a.Date, a.Status, a.Remark); err != nil {
			tx.Rollback(); return err
		}
	}
	return tx.Commit()
}

func (r *AttendanceRepository) ListByClassDate(classID int64, date string) ([]models.Attendance, error) {
	var list []models.Attendance
	return list, r.db.Select(&list, `SELECT * FROM attendance WHERE class_id=? AND date=?`, classID, date)
}

func (r *AttendanceRepository) ListByStudent(studentID, termID int64) ([]models.Attendance, error) {
	var list []models.Attendance
	return list, r.db.Select(&list, `SELECT * FROM attendance WHERE student_id=? AND term_id=? ORDER BY date`, studentID, termID)
}

func (r *AttendanceRepository) SummaryByClass(classID, termID int64) ([]models.AttendanceSummary, error) {
	var list []models.AttendanceSummary
	q := `
		SELECT a.student_id, s.name AS student_name,
		       COUNT(*) AS total,
		       SUM(a.status='present') AS present,
		       SUM(a.status='absent') AS absent,
		       SUM(a.status='late') AS late,
		       ROUND(SUM(a.status='present')/COUNT(*)*100,2) AS percent
		FROM attendance a
		JOIN students s ON s.id=a.student_id
		WHERE a.class_id=? AND a.term_id=?
		GROUP BY a.student_id, s.name`
	return list, r.db.Select(&list, q, classID, termID)
}
