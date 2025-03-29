package user

import (
	"bytes"
	"context"
	"time"

	"github.com/nilhiu/srleaderboard/internal/db"
	"github.com/nilhiu/srleaderboard/internal/db/dto"
	"github.com/nilhiu/srleaderboard/internal/db/models"
)

// Service provides functions to interact with the user database.
type Service interface {
	// Register registers an user, returning a JWT token.
	Register(name, email, password string) (string, error)

	// Login authenticates an user, returning a JWT token.
	Login(username, password string) (string, error)

	// GetUserID returnes the ID of the user with the given username
	GetUserID(username string) (string, error)

	// GetUserRuns returns the runs of the given user of size `limit` from `offset`.
	GetUserRuns(userID string, offset int, limit int) ([]models.Run, error)

	// GetBestRuns returns the best (lowest completion time) runs for each user.
	GetBestRuns() ([]dto.Run, error)

	// GetUserRunCount returns the amount of runs the given user has.
	GetUserRunCount(userID string) (int, error)

	// AddRun adds a run to the given user with the given duration/completion time.
	AddRun(userID string, dur time.Duration) (models.Run, error)
}

type service struct {
	ctx    context.Context
	userDB *db.UserDB
}

func New(ctx context.Context, userDB *db.UserDB) Service {
	return &service{
		ctx:    ctx,
		userDB: userDB,
	}
}

func (s *service) Register(username, email, password string) (string, error) {
	if _, err := s.userDB.GetUser(username); err == nil {
		return "", ErrRegisterUserExists
	}

	err := s.userDB.AddUser(username, email, HashPassword(password))
	if err != nil {
		return "", ErrRegisterInsertFailed
	}

	t, err := newJWT(username)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *service) Login(username, password string) (string, error) {
	usr, err := s.userDB.GetUser(username)
	if err != nil {
		return "", err
	}

	givenHash := HashPassword(password)
	if !bytes.Equal(usr.Password, givenHash) {
		return "", ErrLoginIncorrectPassword
	}

	tok, err := newJWT(username)
	if err != nil {
		return "", err
	}

	return tok, nil
}

func (s *service) GetUserID(username string) (string, error) {
	return s.userDB.GetUserID(username)
}

func (s *service) GetUserRuns(userID string, offset int, limit int) ([]models.Run, error) {
	return s.userDB.GetUserRuns(userID, offset, limit)
}

func (s *service) GetUserRunCount(userID string) (int, error) {
	c, err := s.userDB.GetUserRunCount(userID)
	return int(c), err
}

func (s *service) GetBestRuns() ([]dto.Run, error) {
	return s.userDB.GetBestRuns()
}

func (s *service) AddRun(userID string, dur time.Duration) (models.Run, error) {
	run := models.Run{
		UserID:         userID,
		CompletionTime: dur.Nanoseconds(),
		CreatedAt:      time.Now(),
	}
	err := s.userDB.AddRun(run)

	return run, err
}
