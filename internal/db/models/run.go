package models

import "time"

// A Run represents the internal structure of a run in the SQL database.
type Run struct {
	ID             string    `db:"id"              json:"id"`
	UserID         string    `db:"user_id"         json:"-"`
	CompletionTime int64     `db:"completion_time" json:"completion_time"`
	CreatedAt      time.Time `db:"created_at"      json:"created_at"`
}
