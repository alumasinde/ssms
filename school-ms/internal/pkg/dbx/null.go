package dbx

import "database/sql"

func String(v sql.NullString) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}

func Int(v sql.NullInt64) *int64 {
	if !v.Valid {
		return nil
	}
	return &v.Int64
}