package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nilhiu/srleaderboard/internal/db/dto"
	"github.com/redis/go-redis/v9"
)

// The key of the sorted set used for the leaderboard
const redisLeaderboardKeyName = "leaderboard"

// LeaderboardDB provides methods for interacting with leaderboard data stored
// inside of a Redis sorted set.
type LeaderboardDB struct {
	ctx context.Context
	db  *redis.Client
}

// NewLeaderboardDB connects to a running Redis instance. Requires the following
// environmental variables to be set: `REDIS_HOST`, `REDIS_PORT`, and `REDIS_PASSWORD`.
func NewLeaderboardDB(ctx context.Context) (*LeaderboardDB, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		DB:       0,
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	if db == nil {
		return nil, ErrLeaderboardDBRedisFailed
	}

	return &LeaderboardDB{
		ctx: ctx,
		db:  db,
	}, nil
}

// GetRunCount returns the amount of entries there are in the leaderboard.
func (l *LeaderboardDB) GetRunCount() (int64, error) {
	return l.db.ZCard(l.ctx, redisLeaderboardKeyName).Result()
}

// GetRuns returns a list of runs of size `limit` from an `offset`.
func (l *LeaderboardDB) GetRuns(offset int64, limit int64) ([]dto.Run, error) {
	end := offset + limit - 1
	zs, err := l.db.ZRangeWithScores(l.ctx, redisLeaderboardKeyName, offset, end).Result()
	if err != nil {
		return nil, err
	}

	runs := make([]dto.Run, 0, len(zs))
	for _, z := range zs {
		runs = append(runs, dto.Run{
			Username:       z.Member.(string),
			CompletionTime: int64(z.Score),
		})
	}

	return runs, nil
}

// GetRank returns the rank, or placement, of the given user.
func (l *LeaderboardDB) GetRank(username string) (int64, error) {
	return l.db.ZRank(l.ctx, redisLeaderboardKeyName, username).Result()
}

// AddRun adds a run to the leaderboard. Returns an error only if the given
// [time.Duration] is too small.
func (l *LeaderboardDB) AddRun(username string, dur time.Duration) error {
	// (1 << 53) is the maximal value that can fit into the mantisa
	// of a 64-bit floating point number (in this case float64). Even
	// though larger numbers can be represented with the help of the
	// exponent, this has the risk of not preserving the original time.
	if dur.Nanoseconds() > (1 << 53) {
		return ErrLeaderboardDBRunDurationTooBig
	}

	_, err := l.db.ZAddLT(l.ctx, redisLeaderboardKeyName, redis.Z{
		Score:  float64(dur.Nanoseconds()),
		Member: username,
	}).Result()

	return err
}
