package repositories

import (
	models "school-ms/internal/Modules/Teachers/Models"

	"github.com/jmoiron/sqlx"
)

type TeacherRepository struct{ db *sqlx.DB }

func NewTeacherRepository(db *sqlx.DB) *TeacherRepository { return &TeacherRepository{db: db} }

func (r *TeacherRepository) Create(t *models.Teacher) error {
	q := `INSERT INTO teachers (user_id,school_id,employee_no,phone,gender,dob,qualification,photo_url) VALUES (?,?,?,?,?,?,?,?)`
	res, err := r.db.Exec(q, t.UserID, t.SchoolID, t.EmployeeNo, t.Phone, t.Gender, t.DOB, t.Qualification, t.PhotoURL)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	t.ID = id
	return nil
}

func (r *TeacherRepository) FindByID(id int64) (*models.TeacherDetail, error) {
	var t models.TeacherDetail
	err := r.db.Get(&t, `
		SELECT t.*, u.name, u.email
		FROM teachers t JOIN users u ON u.id = t.user_id
		WHERE t.id=?`, id)
	return &t, err
}

func (r *TeacherRepository) ListBySchool(schoolID int64) ([]models.TeacherDetail, error) {
	var list []models.TeacherDetail
	err := r.db.Select(&list, `
		SELECT t.*, u.name, u.email
		FROM teachers t JOIN users u ON u.id = t.user_id
		WHERE t.school_id=? ORDER BY u.name`, schoolID)
	return list, err
}

func (r *TeacherRepository) Update(t *models.Teacher) error {
	_, err := r.db.Exec(
		`UPDATE teachers SET phone=?,gender=?,dob=?,qualification=?,photo_url=? WHERE id=?`,
		t.Phone, t.Gender, t.DOB, t.Qualification, t.PhotoURL, t.ID,
	)
	return err
}

func (r *TeacherRepository) AssignSubject(ts *models.TeacherSubject) error {
	_, err := r.db.Exec(
		`INSERT IGNORE INTO teacher_subjects (teacher_id,subject_id,class_id) VALUES (?,?,?)`,
		ts.TeacherID, ts.SubjectID, ts.ClassID,
	)
	return err
}

func (r *TeacherRepository) GetSubjectsByTeacher(teacherID int64) ([]models.TeacherSubject, error) {
	var list []models.TeacherSubject
	err := r.db.Select(&list, `SELECT * FROM teacher_subjects WHERE teacher_id=?`, teacherID)
	return list, err
}
