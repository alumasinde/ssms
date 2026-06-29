package repositories

import (
	models "school-ms/internal/Modules/Schools/Models"

	"github.com/jmoiron/sqlx"
)

type SchoolRepository struct {
	db *sqlx.DB
}

func NewSchoolRepository(db *sqlx.DB) *SchoolRepository {
	return &SchoolRepository{db: db}
}

// --- Schools ---

func (r *SchoolRepository) CreateSchool(s *models.School) error {
	q := `INSERT INTO schools (tenant_id,name,code,address,phone,email,logo_url) VALUES (?,?,?,?,?,?,?)`
	res, err := r.db.Exec(q, s.TenantID, s.Name, s.Code, s.Address, s.Phone, s.Email, s.LogoURL)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	s.ID = id
	return nil
}

func (r *SchoolRepository) FindSchoolByID(id int64) (*models.School, error) {
	var s models.School
	return &s, r.db.Get(&s, `SELECT * FROM schools WHERE id=?`, id)
}

func (r *SchoolRepository) ListSchoolsByTenant(tenantID int64) ([]models.School, error) {
	var list []models.School
	return list, r.db.Select(&list, `SELECT * FROM schools WHERE tenant_id=? ORDER BY name`, tenantID)
}

func (r *SchoolRepository) UpdateSchool(s *models.School) error {
	_, err := r.db.Exec(
		`UPDATE schools SET name=?,code=?,address=?,phone=?,email=?,logo_url=? WHERE id=?`,
		s.Name, s.Code, s.Address, s.Phone, s.Email, s.LogoURL, s.ID,
	)
	return err
}

// --- Academic Years ---

func (r *SchoolRepository) CreateAcademicYear(ay *models.AcademicYear) error {
	q := `INSERT INTO academic_years (school_id,name,start_date,end_date,is_current) VALUES (?,?,?,?,?)`
	res, err := r.db.Exec(q, ay.SchoolID, ay.Name, ay.StartDate, ay.EndDate, ay.IsCurrent)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	ay.ID = id
	return nil
}

func (r *SchoolRepository) ListAcademicYears(schoolID int64) ([]models.AcademicYear, error) {
	var list []models.AcademicYear
	return list, r.db.Select(&list, `SELECT * FROM academic_years WHERE school_id=? ORDER BY start_date DESC`, schoolID)
}

func (r *SchoolRepository) GetCurrentAcademicYear(schoolID int64) (*models.AcademicYear, error) {
	var ay models.AcademicYear
	return &ay, r.db.Get(&ay, `SELECT * FROM academic_years WHERE school_id=? AND is_current=1 LIMIT 1`, schoolID)
}

// --- Terms ---

func (r *SchoolRepository) CreateTerm(t *models.Term) error {
	q := `INSERT INTO terms (academic_year_id,name,start_date,end_date,is_current) VALUES (?,?,?,?,?)`
	res, err := r.db.Exec(q, t.AcademicYearID, t.Name, t.StartDate, t.EndDate, t.IsCurrent)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	t.ID = id
	return nil
}

func (r *SchoolRepository) ListTerms(academicYearID int64) ([]models.Term, error) {
	var list []models.Term
	return list, r.db.Select(&list, `SELECT * FROM terms WHERE academic_year_id=? ORDER BY start_date`, academicYearID)
}

func (r *SchoolRepository) GetCurrentTerm(schoolID int64) (*models.Term, error) {
	var t models.Term
	err := r.db.Get(&t, `
		SELECT t.* FROM terms t
		JOIN academic_years ay ON ay.id = t.academic_year_id
		WHERE ay.school_id=? AND t.is_current=1 LIMIT 1`, schoolID)
	return &t, err
}
