package models

// A User represents the internal structure of a user in the SQL database.
type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password []byte `db:"password"`
}
