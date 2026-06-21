package repositories

import (
	"context"
	"time"

	models "school-ms/internal/Modules/Classes/Models"

	"github.com/jmoiron/sqlx"
)

const classTimeout = 5 * time.Second

// explicit column list avoids scanning deleted_at / deleted_by into the model
const classCols = `id, school_id, name, level, stream, created_at, updated_at`

type ClassRepository struct{ db *sqlx.DB }

func NewClassRepository(db *sqlx.DB) *ClassRepository { return &ClassRepository{db: db} }

func (r *ClassRepository) Create(c *models.Class) error {
	ctx, cancel := context.WithTimeout(context.Background(), classTimeout)
	defer cancel()

	// Store NULL for empty stream, not an empty string
	var stream *string
	if c.Stream != nil && *c.Stream != "" {
		stream = c.Stream
	}

	res, err := r.db.ExecContext(ctx, `
		INSERT INTO classes (school_id, name, level, stream)
		VALUES (?, ?, ?, ?)
	`, c.SchoolID, c.Name, c.Level, stream)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	c.ID = id
	return nil
}

func (r *ClassRepository) FindByID(id int64) (*models.Class, error) {
	ctx, cancel := context.WithTimeout(context.Background(), classTimeout)
	defer cancel()

	var c models.Class
	err := r.db.GetContext(ctx, &c, `
		SELECT `+classCols+`
		FROM classes
		WHERE id = ?
		  AND deleted_at IS NULL
	`, id)
	return &c, err
}

func (r *ClassRepository) ListBySchool(schoolID int64) ([]models.Class, error) {
	ctx, cancel := context.WithTimeout(context.Background(), classTimeout)
	defer cancel()

	var list []models.Class
	err := r.db.SelectContext(ctx, &list, `
		SELECT `+classCols+`
		FROM classes
		WHERE school_id  = ?
		  AND deleted_at IS NULL
		ORDER BY level, name
	`, schoolID)
	return list, err
}

func (r *ClassRepository) Update(c *models.Class) error {
	ctx, cancel := context.WithTimeout(context.Background(), classTimeout)
	defer cancel()

	var stream *string
	if c.Stream != nil && *c.Stream != "" {
		stream = c.Stream
	}

	_, err := r.db.ExecContext(ctx, `
		UPDATE classes
		SET name = ?, level = ?, stream = ?
		WHERE id         = ?
		  AND deleted_at IS NULL
	`, c.Name, c.Level, stream, c.ID)
	return err
}

// SoftDelete marks the class as deleted without removing the row.
// Uses deleted_by to record who triggered it.
func (r *ClassRepository) SoftDelete(id, deletedBy int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), classTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `
		UPDATE classes
		SET deleted_at = NOW(),
		    deleted_by = ?
		WHERE id         = ?
		  AND deleted_at IS NULL
	`, deletedBy, id)
	return err
}
