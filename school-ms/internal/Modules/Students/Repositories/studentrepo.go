package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	models "school-ms/internal/Modules/Students/Models"

	"github.com/jmoiron/sqlx"
)

const sTimeout = 5 * time.Second

// studentCols covers ALL columns defined in models.Student.
// Includes medical_notes, left_date, left_reason — missing in original.
const studentCols = `
	s.id, s.school_id, s.class_id, s.admission_no,
	s.first_name, s.middle_name, s.last_name,
	s.gender, s.dob, s.nationality, s.national_id,
	s.religion, s.blood_group, s.address, s.medical_notes,
	s.photo_url, s.is_active, s.enrolled_at,
	s.left_date, s.left_reason`

type StudentRepository struct{ db *sqlx.DB }

func NewStudentRepository(db *sqlx.DB) *StudentRepository { return &StudentRepository{db: db} }

func (r *StudentRepository) Create(s *models.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	res, err := r.db.ExecContext(ctx, `
		INSERT INTO students
			(school_id,class_id,admission_no,first_name,middle_name,last_name,
			 gender,dob,nationality,national_id,religion,blood_group,
			 address,medical_notes,photo_url,is_active)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)
	`, s.SchoolID, s.ClassID, s.AdmissionNo,
		s.FirstName, s.MiddleName, s.LastName,
		s.Gender, s.DOB, s.Nationality, s.NationalID,
		s.Religion, s.BloodGroup, s.Address, s.MedicalNotes, s.PhotoURL)
	if err != nil { return err }
	id, _ := res.LastInsertId(); s.ID = id; return nil
}

func (r *StudentRepository) FindByID(id int64) (*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var s models.Student
	err := r.db.GetContext(ctx, &s, `
		SELECT `+studentCols+`
		FROM students s
		WHERE s.id=? AND s.deleted_at IS NULL
	`, id)
	return &s, err
}

func (r *StudentRepository) FindByAdmissionNo(no string, schoolID int64) (*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var s models.Student
	err := r.db.GetContext(ctx, &s, `
		SELECT `+studentCols+`
		FROM students s
		WHERE s.admission_no=? AND s.school_id=? AND s.deleted_at IS NULL
	`, no, schoolID)
	return &s, err
}

func (r *StudentRepository) ListBySchool(schoolID int64, page, perPage int) ([]models.Student, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var total int64
	r.db.GetContext(ctx, &total, `
		SELECT COUNT(*) FROM students WHERE school_id=? AND is_active=1 AND deleted_at IS NULL
	`, schoolID)
	var list []models.Student
	offset := (page - 1) * perPage
	err := r.db.SelectContext(ctx, &list, `
		SELECT `+studentCols+`
		FROM students s
		WHERE s.school_id=? AND s.is_active=1 AND s.deleted_at IS NULL
		ORDER BY s.first_name, s.last_name
		LIMIT ? OFFSET ?
	`, schoolID, perPage, offset)
	return list, total, err
}

func (r *StudentRepository) ListByClass(classID int64) ([]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var list []models.Student
	err := r.db.SelectContext(ctx, &list, `
		SELECT `+studentCols+`
		FROM students s
		WHERE s.class_id=? AND s.is_active=1 AND s.deleted_at IS NULL
		ORDER BY s.first_name, s.last_name
	`, classID)
	return list, err
}

// Search returns students whose name or admission_no contains q (case-insensitive).
func (r *StudentRepository) Search(schoolID int64, q string) ([]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	like := "%" + q + "%"
	var list []models.Student
	err := r.db.SelectContext(ctx, &list, `
		SELECT `+studentCols+`
		FROM students s
		WHERE s.school_id=? AND s.is_active=1 AND s.deleted_at IS NULL
		  AND (s.first_name LIKE ? OR s.last_name LIKE ?
		       OR s.admission_no LIKE ?
		       OR CONCAT(s.first_name,' ',s.last_name) LIKE ?)
		ORDER BY s.first_name, s.last_name
		LIMIT 50
	`, schoolID, like, like, like, like)
	return list, err
}

func (r *StudentRepository) FindByParentUser(userID int64) ([]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var list []models.Student
	err := r.db.SelectContext(ctx, &list, `
		SELECT DISTINCT `+studentCols+`
		FROM students s
		JOIN parent_student ps ON ps.student_id=s.id
		JOIN parents p ON p.id=ps.parent_id
		WHERE p.user_id=? AND s.is_active=1 AND s.deleted_at IS NULL
		ORDER BY s.first_name, s.last_name
	`, userID)
	return list, err
}

func (r *StudentRepository) FindByTeacherUser(userID int64) ([]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var list []models.Student
	err := r.db.SelectContext(ctx, &list, `
		SELECT DISTINCT `+studentCols+`
		FROM students s
		JOIN teacher_subjects ts ON ts.class_id=s.class_id
		JOIN teachers t ON t.id=ts.teacher_id
		WHERE t.user_id=? AND s.is_active=1 AND s.deleted_at IS NULL
		ORDER BY s.first_name, s.last_name
	`, userID)
	return list, err
}

func (r *StudentRepository) IsParentOfStudent(userID, studentID int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var count int
	err := r.db.QueryRowContext(ctx, `
		SELECT COUNT(*) FROM parent_student ps
		JOIN parents p ON p.id=ps.parent_id
		WHERE p.user_id=? AND ps.student_id=?
	`, userID, studentID).Scan(&count)
	return count > 0, err
}

func (r *StudentRepository) IsTeacherOfStudent(userID, studentID int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var count int
	err := r.db.QueryRowContext(ctx, `
		SELECT COUNT(*) FROM students s
		JOIN teacher_subjects ts ON ts.class_id=s.class_id
		JOIN teachers t ON t.id=ts.teacher_id
		WHERE t.user_id=? AND s.id=?
	`, userID, studentID).Scan(&count)
	return count > 0, err
}

func (r *StudentRepository) Update(s *models.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		UPDATE students SET
			class_id=?, first_name=?, middle_name=?, last_name=?,
			gender=?, dob=?, nationality=?, national_id=?,
			religion=?, blood_group=?, address=?, medical_notes=?, photo_url=?
		WHERE id=? AND deleted_at IS NULL
	`, s.ClassID, s.FirstName, s.MiddleName, s.LastName,
		s.Gender, s.DOB, s.Nationality, s.NationalID,
		s.Religion, s.BloodGroup, s.Address, s.MedicalNotes, s.PhotoURL, s.ID)
	return err
}

// SoftDelete marks the student as deleted and records who deleted.
func (r *StudentRepository) SoftDelete(id, deletedBy int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `
		UPDATE students
		SET is_active=0, deleted_at=NOW(), deleted_by=?
		WHERE id=? AND deleted_at IS NULL
	`, deletedBy, id)
	return err
}

// PromoteToClass moves students to a new class AND writes student_class_history.
func (r *StudentRepository) PromoteToClass(
	fromClassID, toClassID, academicYearID, actorID int64,
	studentIDs []int64,
) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil { return 0, err }

	var ids []int64
	if len(studentIDs) == 0 {
		// Promote entire class
		if err := tx.Select(&ids, `
			SELECT id FROM students
			WHERE class_id=? AND is_active=1 AND deleted_at IS NULL
		`, fromClassID); err != nil {
			tx.Rollback(); return 0, err
		}
	} else {
		ids = studentIDs
	}

	count := 0
	for _, sid := range ids {
		res, err := tx.Exec(`
			UPDATE students SET class_id=?
			WHERE id=? AND class_id=? AND is_active=1 AND deleted_at IS NULL
		`, toClassID, sid, fromClassID)
		if err != nil { tx.Rollback(); return count, err }
		affected, _ := res.RowsAffected()
		if affected == 0 { continue } // student wasn't in fromClass

		// Write history row
		if _, err := tx.Exec(`
			INSERT INTO student_class_history (student_id,class_id,academic_year_id,promoted_by)
			VALUES (?,?,?,?)
		`, sid, toClassID, academicYearID, actorID); err != nil {
			tx.Rollback(); return count, fmt.Errorf("history insert failed: %w", err)
		}
		count++
	}
	return count, tx.Commit()
}

// ParentSummary is the enriched parent projection for GET /students/{id}/parents.
// users has first_name+last_name (no `name` column).
type ParentSummary struct {
	ID           int64  `db:"id"           json:"id"`
	FirstName    string `db:"first_name"   json:"first_name"`
	LastName     string `db:"last_name"    json:"last_name"`
	Email        string `db:"email"        json:"email"`
	Phone        string `db:"phone"        json:"phone"`
	Relationship string `db:"relationship" json:"relationship"`
}

func (p *ParentSummary) FullName() string {
	return strings.TrimSpace(p.FirstName + " " + p.LastName)
}

func (r *StudentRepository) GetParentsByStudent(studentID int64) ([]ParentSummary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), sTimeout)
	defer cancel()
	var list []ParentSummary
	err := r.db.SelectContext(ctx, &list, `
		SELECT p.id,
		       u.first_name, u.last_name,
		       u.email,
		       COALESCE(p.phone,'') AS phone,
		       ps.relationship
		FROM parent_student ps
		JOIN parents p ON p.id=ps.parent_id
		JOIN users   u ON u.id=p.user_id
		WHERE ps.student_id=?
		ORDER BY u.first_name, u.last_name
	`, studentID)
	return list, err
}