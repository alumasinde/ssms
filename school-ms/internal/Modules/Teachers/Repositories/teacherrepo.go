package repositories

import (
	"context"
	"time"

	models "school-ms/internal/Modules/Teachers/Models"

	"github.com/jmoiron/sqlx"
)

const tTimeout = 5 * time.Second

// teacherCols — explicit list avoids scanning deleted_at/deleted_by into the model.
// users has first_name + last_name (no `name` column).
const teacherCols = `
	t.id, t.user_id, t.school_id, t.employee_no,
	t.phone, t.gender, t.dob, t.qualification, t.photo_url,
	t.tsc_no, t.specialization, t.hire_date,
	t.employment_type, t.is_class_teacher, t.national_id, t.address,
	t.is_active, t.created_at, t.updated_at,
	u.first_name, u.last_name, u.email`

type TeacherRepository struct{ db *sqlx.DB }

func NewTeacherRepository(db *sqlx.DB) *TeacherRepository { return &TeacherRepository{db: db} }

func (r *TeacherRepository) Create(t *models.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	if t.EmploymentType == "" { t.EmploymentType = "permanent" }
	res, err := r.db.ExecContext(ctx, `
		INSERT INTO teachers
			(user_id,school_id,employee_no,phone,gender,dob,
			 tsc_no,specialization,hire_date,employment_type,
			 qualification,national_id,address,photo_url)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`, t.UserID, t.SchoolID, t.EmployeeNo, t.Phone, t.Gender, t.DOB,
		t.TSCNo, t.Specialization, t.HireDate, t.EmploymentType,
		t.Qualification, t.NationalID, t.Address, t.PhotoURL)
	if err != nil { return err }
	id, _ := res.LastInsertId(); t.ID = id; return nil
}

func (r *TeacherRepository) FindByID(id int64) (*models.TeacherDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	var t models.TeacherDetail
	err := r.db.GetContext(ctx, &t, `
		SELECT `+teacherCols+`
		FROM teachers t JOIN users u ON u.id=t.user_id
		WHERE t.id=? AND t.deleted_at IS NULL
	`, id)
	return &t, err
}

func (r *TeacherRepository) ListBySchool(schoolID int64) ([]models.TeacherDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	var list []models.TeacherDetail
	err := r.db.SelectContext(ctx, &list, `
		SELECT `+teacherCols+`
		FROM teachers t JOIN users u ON u.id=t.user_id
		WHERE t.school_id=? AND t.is_active=1 AND t.deleted_at IS NULL
		ORDER BY u.first_name, u.last_name
	`, schoolID)
	return list, err
}

func (r *TeacherRepository) Update(t *models.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		UPDATE teachers SET
			phone=?, gender=?, dob=?, tsc_no=?, hire_date=?,
			qualification=?, specialization=?, national_id=?, address=?,
			photo_url=?, is_class_teacher=?, employment_type=?
		WHERE id=? AND deleted_at IS NULL
	`, t.Phone, t.Gender, t.DOB, t.TSCNo, t.HireDate,
		t.Qualification, t.Specialization, t.NationalID, t.Address,
		t.PhotoURL, t.IsClassTeacher, t.EmploymentType, t.ID)
	return err
}

// SoftDelete records who deleted and when; preserves FK referential integrity.
func (r *TeacherRepository) SoftDelete(id, deletedBy int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		UPDATE teachers SET is_active=0, deleted_at=NOW(), deleted_by=?
		WHERE id=? AND deleted_at IS NULL
	`, deletedBy, id)
	return err
}

// AssignSubject — teacher_subjects.school_id is NOT NULL (FK to schools).
func (r *TeacherRepository) AssignSubject(ts *models.TeacherSubject) error {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		INSERT IGNORE INTO teacher_subjects (teacher_id,subject_id,class_id,school_id)
		VALUES (?,?,?,?)
	`, ts.TeacherID, ts.SubjectID, ts.ClassID, ts.SchoolID)
	return err
}

func (r *TeacherRepository) GetSubjectsByTeacher(teacherID int64) ([]models.TeacherSubjectDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	var list []models.TeacherSubjectDetail
	err := r.db.SelectContext(ctx, &list, `
		SELECT ts.id, ts.teacher_id, ts.subject_id, ts.class_id, ts.school_id,
		       s.name AS subject_name, s.code AS subject_code, c.name AS class_name
		FROM teacher_subjects ts
		JOIN subjects s ON s.id=ts.subject_id
		JOIN classes  c ON c.id=ts.class_id
		WHERE ts.teacher_id=?
		  AND (s.deleted_at IS NULL OR s.deleted_at > NOW())
		  AND (c.deleted_at IS NULL OR c.deleted_at > NOW())
		ORDER BY c.name, s.name
	`, teacherID)
	return list, err
}

func (r *TeacherRepository) RemoveSubject(teacherID, subjectID, classID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), tTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM teacher_subjects WHERE teacher_id=? AND subject_id=? AND class_id=?
	`, teacherID, subjectID, classID)
	return err
}