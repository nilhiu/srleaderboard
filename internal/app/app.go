package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/nilhiu/srleaderboard/internal/db"
	"github.com/nilhiu/srleaderboard/internal/service/leaderboard"
	"github.com/nilhiu/srleaderboard/internal/service/user"
)

type App struct {
	ctx            context.Context
	logger         *slog.Logger
	router         *http.ServeMux
	userSvc        user.Service
	leaderboardSvc leaderboard.Service
}

func New(ctx context.Context, logger *slog.Logger) (*App, error) {
	router := http.NewServeMux()

	userDB, err := db.NewUserDB(ctx)
	if err != nil {
		return nil, err
	}

	leaderboardDB, err := db.NewLeaderboardDB(ctx)
	if err != nil {
		return nil, err
	}

	userSvc := user.New(ctx, userDB)
	leaderboardSvc := leaderboard.New(leaderboardDB)

	if err := leaderboardSvc.Initialize(userSvc); err != nil {
		return nil, err
	}

	app := &App{
		ctx:            ctx,
		logger:         logger,
		router:         router,
		userSvc:        userSvc,
		leaderboardSvc: leaderboardSvc,
	}

	registerRoutes(ctx, app)

	return app, nil
}

func (a *App) Run() {
	a.logger.Info("server started", "port", "80")
	if err := http.ListenAndServe(":80", a.router); err != nil {
		a.logger.ErrorContext(a.ctx, "app stopped", "err", err.Error())
		os.Exit(1)
	}
}
