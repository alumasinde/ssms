package models

type Parent struct {
	ID         int64  `db:"id" json:"id"`
	UserID     int64  `db:"user_id" json:"user_id"`
	SchoolID   int64  `db:"school_id" json:"school_id"`
	Phone      string `db:"phone" json:"phone"`
	Occupation string `db:"occupation" json:"occupation"`
	Address    string `db:"address" json:"address"`
}

type ParentStudent struct {
	ID           int64  `db:"id" json:"id"`
	ParentID     int64  `db:"parent_id" json:"parent_id"`
	StudentID    int64  `db:"student_id" json:"student_id"`
	Relationship string `db:"relationship" json:"relationship"`
}

type ParentDetail struct {
	Parent

	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}