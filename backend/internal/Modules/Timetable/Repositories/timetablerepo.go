package repositories

import (
	models "school-ms/internal/Modules/Timetable/Models"
	"github.com/jmoiron/sqlx"
)

type TimetableRepository struct{ db *sqlx.DB }
func NewTimetableRepository(db *sqlx.DB) *TimetableRepository { return &TimetableRepository{db: db} }

const ttJoin = ` FROM timetable_slots ts
	JOIN classes  c   ON c.id   = ts.class_id
	JOIN subjects sub ON sub.id = ts.subject_id
	JOIN teachers t   ON t.id   = ts.teacher_id
	JOIN users    u   ON u.id   = t.user_id`

const ttCols = `ts.id,ts.school_id,ts.class_id,ts.subject_id,ts.teacher_id,
	ts.term_id,ts.day_of_week,ts.start_time,ts.end_time,
	COALESCE(ts.room,'') AS room,
	c.name AS class_name, sub.name AS subject_name, sub.code AS subject_code, u.name AS teacher_name`

func (r *TimetableRepository) Create(s *models.TimetableSlot) error {
	res, err := r.db.Exec(
		`INSERT INTO timetable_slots (school_id,class_id,subject_id,teacher_id,term_id,day_of_week,start_time,end_time,room) VALUES (?,?,?,?,?,?,?,?,?)`,
		s.SchoolID,s.ClassID,s.SubjectID,s.TeacherID,s.TermID,s.DayOfWeek,s.StartTime,s.EndTime,s.Room)
	if err != nil { return err }
	id,_ := res.LastInsertId(); s.ID = id; return nil
}
func (r *TimetableRepository) Update(s *models.TimetableSlot) error {
	_, err := r.db.Exec(`UPDATE timetable_slots SET subject_id=?,teacher_id=?,day_of_week=?,start_time=?,end_time=?,room=? WHERE id=?`,
		s.SubjectID,s.TeacherID,s.DayOfWeek,s.StartTime,s.EndTime,s.Room,s.ID)
	return err
}
func (r *TimetableRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM timetable_slots WHERE id=?`, id); return err
}
func (r *TimetableRepository) ListByClass(classID,termID int64) ([]models.TimetableSlotDetail,error) {
	var list []models.TimetableSlotDetail
	return list, r.db.Select(&list, `SELECT `+ttCols+ttJoin+` WHERE ts.class_id=? AND ts.term_id=? ORDER BY ts.day_of_week,ts.start_time`, classID,termID)
}
func (r *TimetableRepository) ListByTeacher(teacherID,termID int64) ([]models.TimetableSlotDetail,error) {
	var list []models.TimetableSlotDetail
	return list, r.db.Select(&list, `SELECT `+ttCols+ttJoin+` WHERE ts.teacher_id=? AND ts.term_id=? ORDER BY ts.day_of_week,ts.start_time`, teacherID,termID)
}
func (r *TimetableRepository) ListBySchool(schoolID,termID int64) ([]models.TimetableSlotDetail,error) {
	var list []models.TimetableSlotDetail
	return list, r.db.Select(&list, `SELECT `+ttCols+ttJoin+` WHERE ts.school_id=? AND ts.term_id=? ORDER BY ts.class_id,ts.day_of_week,ts.start_time`, schoolID,termID)
}
