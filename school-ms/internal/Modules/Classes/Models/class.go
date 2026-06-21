package models

import "time"

// Class mirrors the classes table in the live schema.
//
// Schema notes:
//   - stream is nullable (varchar(40) DEFAULT NULL) → *string
//   - soft-delete via deleted_at / deleted_by
//   - created_at / updated_at managed by MySQL
type Class struct {
	ID        int64      `db:"id"         json:"id"`
	SchoolID  int64      `db:"school_id"  json:"school_id"`
	Name      string     `db:"name"       json:"name"`
	Level     string     `db:"level"      json:"level"`
	Stream    *string    `db:"stream"     json:"stream,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}
