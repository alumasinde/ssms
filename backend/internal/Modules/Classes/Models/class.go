package models

import "time"

type Class struct {
	ID        int64      `db:"id"         json:"id"`
	SchoolID  int64      `db:"school_id"  json:"school_id"`
	Name      string     `db:"name"       json:"name"`
	Level     string     `db:"level"      json:"level"`
	Stream    *string    `db:"stream"     json:"stream,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}
