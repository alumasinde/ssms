package repositories

import (
	"context"
	"time"

	models "school-ms/internal/Modules/Classes/Models"

	"github.com/jmoiron/sqlx"
)

const csTimeout = 5 * time.Second

type ClassSubjectRepository struct{ db *sqlx.DB }

func NewClassSubjectRepository(db *sqlx.DB) *ClassSubjectRepository {
	return &ClassSubjectRepository{db: db}
}

// AssignSubject inserts or updates a subject assignment for a class.
// school_id is required by the schema (NOT NULL FK).
func (r *ClassSubjectRepository) AssignSubject(
	classID, subjectID, schoolID int64,
	compulsory bool,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), csTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO class_subjects (class_id, subject_id, school_id, is_compulsory)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			is_compulsory = VALUES(is_compulsory)
	`, classID, subjectID, schoolID, compulsory)
	return err
}

// RemoveSubject deletes a subject assignment from a class.
func (r *ClassSubjectRepository) RemoveSubject(classID, subjectID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), csTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `
		DELETE FROM class_subjects
		WHERE class_id = ? AND subject_id = ?
	`, classID, subjectID)
	return err
}

// ListSubjects returns all subjects assigned to a class, enriched with
// subject name and code from the subjects table.
func (r *ClassSubjectRepository) ListSubjects(classID int64) ([]models.ClassSubject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), csTimeout)
	defer cancel()

	var list []models.ClassSubject
	err := r.db.SelectContext(ctx, &list, `
		SELECT
			cs.id,
			cs.class_id,
			cs.subject_id,
			cs.school_id,
			cs.is_compulsory,
			s.name AS subject_name,
			s.code AS subject_code
		FROM class_subjects cs
		JOIN subjects s ON s.id = cs.subject_id
		WHERE cs.class_id = ?
		  AND (s.deleted_at IS NULL OR s.deleted_at > NOW())
		ORDER BY s.name
	`, classID)
	return list, err
}

// ListUnassignedSubjects returns school subjects not yet assigned to this class.
func (r *ClassSubjectRepository) ListUnassignedSubjects(
	classID, schoolID int64,
) ([]models.ClassSubject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), csTimeout)
	defer cancel()

	var list []models.ClassSubject
	err := r.db.SelectContext(ctx, &list, `
		SELECT
			0         AS id,
			?         AS class_id,
			s.id      AS subject_id,
			?         AS school_id,
			1         AS is_compulsory,
			s.name    AS subject_name,
			s.code    AS subject_code
		FROM subjects s
		WHERE s.school_id = ?
		  AND s.is_active = 1
		  AND (s.deleted_at IS NULL OR s.deleted_at > NOW())
		  AND s.id NOT IN (
			SELECT subject_id
			FROM class_subjects
			WHERE class_id = ?
		  )
		ORDER BY s.name
	`, classID, schoolID, schoolID, classID)
	return list, err
}
