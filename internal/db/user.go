package db

import (
	"context"
	"embed"
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/nilhiu/srleaderboard/internal/db/dto"
	"github.com/nilhiu/srleaderboard/internal/db/models"
)

// migrations store all the migration files that are located in the `migrations/` folder
//
//go:embed migrations/*.sql
var migrations embed.FS

// UserDB provides methods for interacting with user data stored inside of a
// PostgreSQL database.
type UserDB struct {
	ctx context.Context
	db  *pgx.Conn
}

// NewLeaderboardDB connects to a running PostgreSQL instance. Requires the
// `POSTGRES_URL` environmental variable to be set.
func NewUserDB(ctx context.Context) (*UserDB, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}

	db := &UserDB{
		ctx: ctx,
		db:  conn,
	}

	if err := db.migrate(); err != nil {
		return nil, err
	}

	return db, nil
}

func (u *UserDB) migrate() error {
	src, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithSourceInstance("iofs", src, os.Getenv("POSTGRES_URL"))
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}

		return err
	}

	return nil
}

// GetUser returns an user with the given name.
func (u *UserDB) GetUser(name string) (models.User, error) {
	var id, email string
	var pass []byte

	err := u.db.QueryRow(u.ctx, "SELECT id, email, password FROM users WHERE name = $1", name).
		Scan(&id, &email, &pass)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:       id,
		Email:    email,
		Name:     name,
		Password: pass,
	}, nil
}

// GetUserID returns the user's ID with the given name.
func (u *UserDB) GetUserID(name string) (string, error) {
	var id string

	err := u.db.QueryRow(u.ctx, "SELECT id FROM users WHERE name = $1", name).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// AddUser adds a user to the database. The `passHash` has to be the return
// value of the [user.HashPassword] function.
func (u *UserDB) AddUser(name, email string, passHash []byte) error {
	_, err := u.db.Exec(
		u.ctx,
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		name,
		email,
		passHash,
	)

	return err
}

// AddRun adds a run to the database.
func (u *UserDB) AddRun(run models.Run) error {
	_, err := u.db.Exec(
		u.ctx,
		"INSERT INTO runs (user_id, completion_time, created_at) VALUES ($1, $2, $3)",
		run.UserID,
		run.CompletionTime,
		run.CreatedAt,
	)

	return err
}

// GetUserRuns returns the user's runs, by the given ID, of size `limit` and from an `offset`.
func (u *UserDB) GetUserRuns(userID string, offset int, limit int) ([]models.Run, error) {
	rows, err := u.db.Query(
		u.ctx,
		"SELECT id, user_id, completion_time, created_at FROM runs WHERE user_id = $1 OFFSET $2 LIMIT $3",
		userID,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	runs, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Run])
	if err != nil {
		return nil, err
	}

	return runs, nil
}

// GetUserRunCount returns the number of runs a user has, with the given ID.
func (u *UserDB) GetUserRunCount(userID string) (int64, error) {
	var count int64
	err := u.db.QueryRow(
		u.ctx,
		"SELECT COUNT(id) FROM runs WHERE user_id = $1",
		userID,
	).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetBestRuns returns the best (lowest completion time) runs of each user.
func (u *UserDB) GetBestRuns() ([]dto.Run, error) {
	rows, err := u.db.Query(
		u.ctx,
		`
SELECT DISTINCT ON (u.id) u.name, r.completion_time
FROM runs r
JOIN users u ON r.user_id = u.id
ORDER BY u.id, r.completion_time ASC
    `,
	)
	if err != nil {
		return nil, err
	}

	runs, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Run])
	if err != nil {
		return nil, err
	}

	return runs, nil
}
