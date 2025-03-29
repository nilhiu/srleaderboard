package leaderboard

import (
	"time"

	"github.com/nilhiu/srleaderboard/internal/db"
	"github.com/nilhiu/srleaderboard/internal/db/dto"
	"github.com/nilhiu/srleaderboard/internal/service/user"
)

// Service provides functions to interact with the leaderboard database.
type Service interface {
	// Initialize initializes the leaderboard database with the data from the
	// user database.
	Initialize(userSvc user.Service) error

	// AddRun adds a run to the leaderboard.
	AddRun(username string, dur time.Duration) error

	// GetRunCount returns the amount of runs currently on the leaderboard.
	GetRunCount() (int, error)

	// GetRuns return runs from the leaderboard.
	GetRuns(offset int, limit int) ([]dto.Run, error)

	// GetRank gets the rank/placement of the given user.
	GetRank(username string) (int64, error)
}

type service struct {
	leaderboardDB *db.LeaderboardDB
}

func New(leaderboardDB *db.LeaderboardDB) Service {
	return &service{
		leaderboardDB: leaderboardDB,
	}
}

func (s *service) Initialize(userSvc user.Service) error {
	runs, err := userSvc.GetBestRuns()
	if err != nil {
		return err
	}

	for _, run := range runs {
		if err := s.leaderboardDB.AddRun(run.Username, time.Duration(run.CompletionTime)); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetRunCount() (int, error) {
	c, err := s.leaderboardDB.GetRunCount()
	return int(c), err
}

func (s *service) GetRuns(offset int, limit int) ([]dto.Run, error) {
	return s.leaderboardDB.GetRuns(int64(offset), int64(limit))
}

func (s *service) GetRank(username string) (int64, error) {
	return s.leaderboardDB.GetRank(username)
}

func (s *service) AddRun(username string, dur time.Duration) error {
	return s.leaderboardDB.AddRun(username, dur)
}
