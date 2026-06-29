package repositories

import (
	models "school-ms/internal/Modules/Terms/Models"
	"github.com/jmoiron/sqlx"
)

type TermRepository struct{ db *sqlx.DB }

func NewTermRepository(db *sqlx.DB) *TermRepository { return &TermRepository{db: db} }

func (r *TermRepository) Create(t *models.Term) error {
	res, err := r.db.Exec(
		`INSERT INTO terms (academic_year_id, name, start_date, end_date, is_current) VALUES (?,?,?,?,?)`,
		t.AcademicYearID, t.Name, t.StartDate, t.EndDate, t.IsCurrent)
	if err != nil { return err }
	id, _ := res.LastInsertId(); t.ID = id; return nil
}

func (r *TermRepository) FindByID(id int64) (*models.Term, error) {
	var t models.Term
	err := r.db.Get(&t, `
		SELECT t.id, t.academic_year_id, ay.school_id, t.name, t.start_date, t.end_date, t.is_current
		FROM terms t JOIN academic_years ay ON ay.id = t.academic_year_id
		WHERE t.id=?`, id)
	return &t, err
}

func (r *TermRepository) ListByAcademicYear(yearID int64) ([]models.Term, error) {
	var list []models.Term
	err := r.db.Select(&list, `
		SELECT t.id, t.academic_year_id, ay.school_id, t.name, t.start_date, t.end_date, t.is_current
		FROM terms t JOIN academic_years ay ON ay.id = t.academic_year_id
		WHERE t.academic_year_id=? ORDER BY t.start_date`, yearID)
	return list, err
}

func (r *TermRepository) ListBySchool(schoolID int64) ([]models.Term, error) {
	var list []models.Term
	err := r.db.Select(&list, `
		SELECT t.id, t.academic_year_id, ay.school_id, t.name, t.start_date, t.end_date, t.is_current
		FROM terms t JOIN academic_years ay ON ay.id = t.academic_year_id
		WHERE ay.school_id=? ORDER BY ay.start_date DESC, t.start_date`, schoolID)
	return list, err
}

func (r *TermRepository) GetCurrent(schoolID int64) (*models.Term, error) {
	var t models.Term
	err := r.db.Get(&t, `
		SELECT t.id, t.academic_year_id, ay.school_id, t.name, t.start_date, t.end_date, t.is_current
		FROM terms t JOIN academic_years ay ON ay.id = t.academic_year_id
		WHERE ay.school_id=? AND t.is_current=1 LIMIT 1`, schoolID)
	return &t, err
}

func (r *TermRepository) SetCurrent(schoolID, termID int64) error {
	tx, err := r.db.Beginx()
	if err != nil { return err }
	tx.Exec(`UPDATE terms t JOIN academic_years ay ON ay.id=t.academic_year_id
		SET t.is_current=0 WHERE ay.school_id=?`, schoolID)
	tx.Exec(`UPDATE terms SET is_current=1 WHERE id=?`, termID)
	return tx.Commit()
}

func (r *TermRepository) Update(t *models.Term) error {
	_, err := r.db.Exec(
		`UPDATE terms SET name=?, start_date=?, end_date=? WHERE id=?`,
		t.Name, t.StartDate, t.EndDate, t.ID)
	return err
}

func (r *TermRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM terms WHERE id=?`, id); return err
}
