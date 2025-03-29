package ajax

import (
	"time"

	"github.com/nilhiu/srleaderboard/internal/db/dto"
	"github.com/nilhiu/srleaderboard/internal/db/models"
)

type AddRunResponse struct {
	Username  string    `json:"username"`
	Placement int64     `json:"placement"`
	Time      Duration  `json:"time"`
	DateAdded time.Time `json:"date_added"`
}

type GetRunsResponse struct {
	Runs       []dto.Run `json:"runs"`
	Amount     int       `json:"amount"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	FullAmount int       `json:"full_amount"`
}

type GetUserRunsResponse struct {
	Runs       []models.Run `json:"runs"`
	Amount     int          `json:"amount"`
	Offset     int          `json:"offset"`
	Limit      int          `json:"limit"`
	FullAmount int          `json:"full_amount"`
}
