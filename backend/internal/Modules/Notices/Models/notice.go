package models

import "time"

type Notice struct {
    ID          int64     `db:"id"          json:"id"`
    SchoolID    int64     `db:"school_id"   json:"school_id"`
    AuthorID    int64     `db:"author_id"   json:"author_id"`
    Title       string    `db:"title"       json:"title"`
    Body        string    `db:"body"        json:"body"`
    Audience    string    `db:"audience"    json:"audience"`
    PublishedAt time.Time `db:"published_at" json:"published_at"`
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`

}
