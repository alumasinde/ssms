package repositories

import (
	models "school-ms/internal/Modules/AcademicYears/Models"
	"github.com/jmoiron/sqlx"
)

type AcademicYearRepository struct{ db *sqlx.DB }

func NewAcademicYearRepository(db *sqlx.DB) *AcademicYearRepository {
	return &AcademicYearRepository{db: db}
}

func (r *AcademicYearRepository) Create(ay *models.AcademicYear) error {
	res, err := r.db.Exec(
		`INSERT INTO academic_years (school_id,name,start_date,end_date,is_current) VALUES (?,?,?,?,?)`,
		ay.SchoolID, ay.Name, ay.StartDate, ay.EndDate, ay.IsCurrent)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	ay.ID = id
	return nil
}

func (r *AcademicYearRepository) FindByID(id int64) (*models.AcademicYear, error) {
	var ay models.AcademicYear
	return &ay, r.db.Get(&ay, `SELECT * FROM academic_years WHERE id=?`, id)
}

func (r *AcademicYearRepository) ListBySchool(schoolID int64) ([]models.AcademicYear, error) {
	var list []models.AcademicYear
	return list, r.db.Select(&list,
		`SELECT * FROM academic_years WHERE school_id=? ORDER BY start_date DESC`, schoolID)
}

func (r *AcademicYearRepository) Update(ay *models.AcademicYear) error {
	_, err := r.db.Exec(
		`UPDATE academic_years SET name=?,start_date=?,end_date=?,is_current=? WHERE id=?`,
		ay.Name, ay.StartDate, ay.EndDate, ay.IsCurrent, ay.ID)
	return err
}

func (r *AcademicYearRepository) SetCurrent(schoolID, yearID int64) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	tx.Exec(`UPDATE academic_years SET is_current=0 WHERE school_id=?`, schoolID)
	tx.Exec(`UPDATE academic_years SET is_current=1 WHERE id=?`, yearID)
	return tx.Commit()
}

func (r *AcademicYearRepository) GetCurrent(schoolID int64) (*models.AcademicYear, error) {
	var ay models.AcademicYear
	return &ay, r.db.Get(&ay,
		`SELECT * FROM academic_years WHERE school_id=? AND is_current=1 LIMIT 1`, schoolID)
}

func (r *AcademicYearRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM academic_years WHERE id=?`, id)
	return err
}
