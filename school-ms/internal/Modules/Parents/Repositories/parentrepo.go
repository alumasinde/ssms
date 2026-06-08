package repositories

import (
	models "school-ms/internal/Modules/Parents/Models"
	"github.com/jmoiron/sqlx"
)

type ParentRepository struct{ db *sqlx.DB }

func NewParentRepository(db *sqlx.DB) *ParentRepository { return &ParentRepository{db: db} }

func (r *ParentRepository) Create(p *models.Parent) error {
	res, err := r.db.Exec(`INSERT INTO parents (user_id,school_id,phone,occupation,address) VALUES (?,?,?,?,?)`,
		p.UserID, p.SchoolID, p.Phone, p.Occupation, p.Address)
	if err != nil { return err }
	id, _ := res.LastInsertId(); p.ID = id; return nil
}

func (r *ParentRepository) FindByID(id int64) (*models.ParentDetail, error) {
	var p models.ParentDetail
	return &p, r.db.Get(&p, `SELECT p.*,u.name,u.email FROM parents p JOIN users u ON u.id=p.user_id WHERE p.id=?`, id)
}

func (r *ParentRepository) ListBySchool(schoolID int64) ([]models.ParentDetail, error) {
	var list []models.ParentDetail
	return list, r.db.Select(&list, `SELECT p.*,u.name,u.email FROM parents p JOIN users u ON u.id=p.user_id WHERE p.school_id=? ORDER BY u.name`, schoolID)
}

func (r *ParentRepository) LinkStudent(ps *models.ParentStudent) error {
	_, err := r.db.Exec(`INSERT IGNORE INTO parent_student (parent_id,student_id,relationship) VALUES (?,?,?)`,
		ps.ParentID, ps.StudentID, ps.Relationship)
	return err
}

func (r *ParentRepository) GetStudentsByParent(parentID int64) ([]int64, error) {
	var ids []int64
	return ids, r.db.Select(&ids, `SELECT student_id FROM parent_student WHERE parent_id=?`, parentID)
}

func (r *ParentRepository) GetParentsByStudent(studentID int64) ([]models.ParentDetail, error) {
	var list []models.ParentDetail
	return list, r.db.Select(&list, `
		SELECT p.*,u.name,u.email FROM parents p
		JOIN users u ON u.id=p.user_id
		JOIN parent_student ps ON ps.parent_id=p.id
		WHERE ps.student_id=?`, studentID)
}
