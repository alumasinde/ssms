package repositories

import (
	models "school-ms/internal/Modules/Discipline/Models"
	"github.com/jmoiron/sqlx"
)

type DisciplineRepository struct{ db *sqlx.DB }
func NewDisciplineRepository(db *sqlx.DB) *DisciplineRepository { return &DisciplineRepository{db:db} }

const dJoin = ` FROM discipline_records dr
	JOIN students s ON s.id = dr.student_id
	JOIN terms    t ON t.id = dr.term_id
	JOIN users    u ON u.id = dr.recorded_by`
const dCols = `dr.id,dr.school_id,dr.student_id,dr.term_id,dr.incident_date,dr.type,
	dr.description,dr.action_taken,dr.recorded_by,dr.created_at,
	CONCAT(s.first_name,' ',s.last_name) AS student_name,
	s.admission_no, t.name AS term_name, u.name AS recorder_name`

func (r *DisciplineRepository) Create(d *models.DisciplineRecord) error {
	res,err := r.db.Exec(
		`INSERT INTO discipline_records (school_id,student_id,term_id,incident_date,type,description,action_taken,recorded_by)
		 VALUES (?,?,?,?,?,?,?,?)`,
		d.SchoolID,d.StudentID,d.TermID,d.IncidentDate,d.Type,d.Description,d.ActionTaken,d.RecordedBy)
	if err != nil { return err }
	id,_ := res.LastInsertId(); d.ID = id; return nil
}
func (r *DisciplineRepository) ListBySchool(schoolID,termID int64) ([]models.DisciplineDetail,error) {
	var list []models.DisciplineDetail
	q := `SELECT `+dCols+dJoin+` WHERE dr.school_id=?`
	args := []interface{}{schoolID}
	if termID > 0 { q += ` AND dr.term_id=?`; args = append(args,termID) }
	q += ` ORDER BY dr.incident_date DESC`
	return list, r.db.Select(&list,q,args...)
}
func (r *DisciplineRepository) ListByStudent(studentID int64) ([]models.DisciplineDetail,error) {
	var list []models.DisciplineDetail
	return list, r.db.Select(&list,`SELECT `+dCols+dJoin+` WHERE dr.student_id=? ORDER BY dr.incident_date DESC`,studentID)
}
func (r *DisciplineRepository) Delete(id int64) error {
	_,err := r.db.Exec(`DELETE FROM discipline_records WHERE id=?`,id); return err
}
