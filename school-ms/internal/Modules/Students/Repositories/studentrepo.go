package repositories

import (
	models "school-ms/internal/Modules/Students/Models"
	"github.com/jmoiron/sqlx"
)

type StudentRepository struct{ db *sqlx.DB }

func NewStudentRepository(db *sqlx.DB) *StudentRepository { return &StudentRepository{db: db} }

func (r *StudentRepository) Create(s *models.Student) error {
	q := `INSERT INTO students (school_id,class_id,admission_no,first_name,middle_name,last_name,gender,dob,photo_url,is_active) VALUES (?,?,?,?,?,?,?,?,?,1)`
	res, err := r.db.Exec(q, s.SchoolID, s.ClassID, s.AdmissionNo, s.FirstName, s.MiddleName, s.LastName, s.Gender, s.DOB, s.PhotoURL)
	if err != nil { return err }
	id, _ := res.LastInsertId(); s.ID = id; return nil
}

func (r *StudentRepository) FindByID(id int64) (*models.Student, error) {
	var s models.Student; return &s, r.db.Get(&s, `SELECT id,school_id,class_id,admission_no,first_name,middle_name,last_name,gender,dob,photo_url,is_active, enrolled_at FROM students WHERE id=?`, id)
}

func (r *StudentRepository) FindByAdmissionNo(no string, schoolID int64) (*models.Student, error) {
	var s models.Student
	return &s, r.db.Get(&s, `SELECT id,school_id,class_id,admission_no,first_name,middle_name,last_name,gender,dob,photo_url,is_active, enrolled_at FROM students WHERE admission_no=? AND school_id=?`, no, schoolID)
}

func (r *StudentRepository) ListBySchool(schoolID int64, page, perPage int) ([]models.Student, int64, error) {
	var list []models.Student
	var total int64
	offset := (page - 1) * perPage
	r.db.Get(&total, `SELECT COUNT(*) FROM students WHERE school_id=? AND is_active=1`, schoolID)
	err := r.db.Select(&list, `SELECT id,school_id,class_id,admission_no,first_name,middle_name,last_name,gender,dob,photo_url,is_active, enrolled_at FROM students WHERE school_id=? AND is_active=1 ORDER BY first_name LIMIT ? OFFSET ?`, schoolID, perPage, offset)
	return list, total, err
}

func (r *StudentRepository) ListByClass(classID int64) ([]models.Student, error) {
	var list []models.Student
	return list, r.db.Select(&list, `SELECT id,school_id,class_id,admission_no,first_name,middle_name,last_name,gender,dob,photo_url,is_active, enrolled_at FROM students WHERE class_id=? AND is_active=1 ORDER BY first_name`, classID)
}
func (r *StudentRepository) FindByParentUser(userID int64) ([]models.Student, error) {
	var list []models.Student
	err := r.db.Select(&list, `
		SELECT DISTINCT s.id, s.school_id, s.class_id, s.admission_no,
		       s.first_name, s.middle_name, s.last_name, s.gender,
		       s.dob, s.photo_url, s.is_active, s.enrolled_at
		FROM students s
		JOIN parent_student ps ON ps.student_id = s.id
		JOIN parents p ON p.id = ps.parent_id
		WHERE p.user_id = ? AND s.is_active = 1
		ORDER BY s.first_name`, userID)
	return list, err
}

func (r *StudentRepository) FindByTeacherUser(userID int64) ([]models.Student, error) {
	var list []models.Student
	err := r.db.Select(&list, `
		SELECT DISTINCT s.id, s.school_id, s.class_id, s.admission_no,
		       s.first_name, s.middle_name, s.last_name, s.gender,
		       s.dob, s.photo_url, s.is_active, s.enrolled_at
		FROM students s
		JOIN teacher_subjects ts ON ts.class_id = s.class_id
		JOIN teachers t ON t.id = ts.teacher_id
		WHERE t.user_id = ? AND s.is_active = 1
		ORDER BY s.first_name`, userID)
	return list, err
}

func (r *StudentRepository) IsParentOfStudent(userID, studentID int64) (bool, error) {
	var count int
	err := r.db.QueryRow(`
		SELECT COUNT(*) FROM parent_student ps
		JOIN parents p ON p.id = ps.parent_id
		WHERE p.user_id = ? AND ps.student_id = ?`, userID, studentID).Scan(&count)
	return count > 0, err
}

func (r *StudentRepository) IsTeacherOfStudent(userID, studentID int64) (bool, error) {
	var count int
	err := r.db.QueryRow(`
		SELECT COUNT(*) FROM students s
		JOIN teacher_subjects ts ON ts.class_id = s.class_id
		JOIN teachers t ON t.id = ts.teacher_id
		WHERE t.user_id = ? AND s.id = ?`, userID, studentID).Scan(&count)
	return count > 0, err
}

func (r *StudentRepository) Update(s *models.Student) error {
	_, err := r.db.Exec(`UPDATE students SET class_id=?,first_name=?,middle_name=?,last_name=?,gender=?,dob=?,photo_url=? WHERE id=?`,
		s.ClassID, s.FirstName, s.MiddleName, s.LastName, s.Gender, s.DOB, s.PhotoURL, s.ID)
	return err
}

func (r *StudentRepository) Deactivate(id int64) error {
	_, err := r.db.Exec(`UPDATE students SET is_active=0 WHERE id=?`, id); return err
}
