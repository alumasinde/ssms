package repositories

import (
	models "school-ms/internal/Modules/Assignments/Models"
	"github.com/jmoiron/sqlx"
)

type AssignmentRepository struct{ db *sqlx.DB }
func NewAssignmentRepository(db *sqlx.DB) *AssignmentRepository { return &AssignmentRepository{db:db} }

const aJoin = ` FROM assignments a
	JOIN classes  c   ON c.id   = a.class_id
	JOIN subjects sub ON sub.id = a.subject_id
	JOIN teachers t   ON t.id   = a.teacher_id
	JOIN users    u   ON u.id   = t.user_id
	JOIN terms    tm  ON tm.id  = a.term_id`
const aCols = `a.id,a.school_id,a.class_id,a.subject_id,a.teacher_id,a.term_id,
	a.title,a.description,a.due_date,a.max_marks,a.created_at,
	c.name AS class_name, sub.name AS subject_name, u.name AS teacher_name, tm.name AS term_name`

func (r *AssignmentRepository) Create(a *models.Assignment) error {
	res,err := r.db.Exec(
		`INSERT INTO assignments (school_id,class_id,subject_id,teacher_id,term_id,title,description,due_date,max_marks) VALUES (?,?,?,?,?,?,?,?,?)`,
		a.SchoolID,a.ClassID,a.SubjectID,a.TeacherID,a.TermID,a.Title,a.Description,a.DueDate,a.MaxMarks)
	if err != nil { return err }
	id,_ := res.LastInsertId(); a.ID = id; return nil
}
func (r *AssignmentRepository) ListByClass(classID,termID int64) ([]models.AssignmentDetail,error) {
	var list []models.AssignmentDetail
	return list, r.db.Select(&list,`SELECT `+aCols+aJoin+` WHERE a.class_id=? AND a.term_id=? ORDER BY a.due_date`,classID,termID)
}
func (r *AssignmentRepository) ListByTeacher(teacherID,termID int64) ([]models.AssignmentDetail,error) {
	var list []models.AssignmentDetail
	return list, r.db.Select(&list,`SELECT `+aCols+aJoin+` WHERE a.teacher_id=? AND a.term_id=? ORDER BY a.due_date`,teacherID,termID)
}
func (r *AssignmentRepository) ListBySchool(schoolID,termID int64) ([]models.AssignmentDetail,error) {
	var list []models.AssignmentDetail
	return list, r.db.Select(&list,`SELECT `+aCols+aJoin+` WHERE a.school_id=? AND a.term_id=? ORDER BY a.due_date`,schoolID,termID)
}
func (r *AssignmentRepository) Delete(id int64) error {
	_,err := r.db.Exec(`DELETE FROM assignments WHERE id=?`,id); return err
}
