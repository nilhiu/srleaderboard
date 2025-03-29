package dto

// A Run represents the data from the SQL database that's used by either for
// sending entries to Redis or for the frontend to display the run.
type Run struct {
	Username       string `db:"name"            json:"username"`
	CompletionTime int64  `db:"completion_time" json:"completion_time"`
}
