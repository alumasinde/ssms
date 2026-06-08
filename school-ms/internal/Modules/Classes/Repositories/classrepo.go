package repositories

import (
	models "school-ms/internal/Modules/Classes/Models"
	"github.com/jmoiron/sqlx"
)

type ClassRepository struct{ db *sqlx.DB }

func NewClassRepository(db *sqlx.DB) *ClassRepository { return &ClassRepository{db: db} }

func (r *ClassRepository) Create(c *models.Class) error {
	res, err := r.db.Exec(`INSERT INTO classes (school_id,name,level,stream) VALUES (?,?,?,?)`, c.SchoolID, c.Name, c.Level, c.Stream)
	if err != nil { return err }
	id, _ := res.LastInsertId(); c.ID = id; return nil
}

func (r *ClassRepository) FindByID(id int64) (*models.Class, error) {
	var c models.Class; return &c, r.db.Get(&c, `SELECT * FROM classes WHERE id=?`, id)
}

func (r *ClassRepository) ListBySchool(schoolID int64) ([]models.Class, error) {
	var list []models.Class
	return list, r.db.Select(&list, `SELECT * FROM classes WHERE school_id=? ORDER BY level,name`, schoolID)
}

func (r *ClassRepository) Update(c *models.Class) error {
	_, err := r.db.Exec(`UPDATE classes SET name=?,level=?,stream=? WHERE id=?`, c.Name, c.Level, c.Stream, c.ID)
	return err
}

func (r *ClassRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM classes WHERE id=?`, id); return err
}
