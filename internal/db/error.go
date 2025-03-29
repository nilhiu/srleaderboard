package db

import "errors"

var (
	ErrLeaderboardDBRunDurationTooBig = errors.New("duration too big")
	ErrLeaderboardDBRedisFailed       = errors.New("redis connection failed")
)
