package repositories

import (
	models "school-ms/internal/Modules/Subjects/Models"
	"github.com/jmoiron/sqlx"
)

type SubjectRepository struct{ db *sqlx.DB }

func NewSubjectRepository(db *sqlx.DB) *SubjectRepository { return &SubjectRepository{db: db} }

func (r *SubjectRepository) Create(s *models.Subject) error {
	res, err := r.db.Exec(`INSERT INTO subjects (school_id,name,code,is_active) VALUES (?,?,?,1)`, s.SchoolID, s.Name, s.Code)
	if err != nil { return err }
	id, _ := res.LastInsertId(); s.ID = id; return nil
}

func (r *SubjectRepository) ListBySchool(schoolID int64) ([]models.Subject, error) {
	var list []models.Subject
	return list, r.db.Select(&list, `SELECT * FROM subjects WHERE school_id=? AND is_active=1 ORDER BY name`, schoolID)
}

func (r *SubjectRepository) FindByID(id int64) (*models.Subject, error) {
	var s models.Subject; return &s, r.db.Get(&s, `SELECT * FROM subjects WHERE id=?`, id)
}

func (r *SubjectRepository) Update(s *models.Subject) error {
	_, err := r.db.Exec(`UPDATE subjects SET name=?,code=? WHERE id=?`, s.Name, s.Code, s.ID); return err
}

func (r *SubjectRepository) Delete(id int64) error {
	_, err := r.db.Exec(`UPDATE subjects SET is_active=0 WHERE id=?`, id); return err
}
