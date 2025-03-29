package app

import (
	"context"
	"embed"
	"net/http"

	"github.com/a-h/templ"
	"github.com/nilhiu/srleaderboard/internal/app/handler"
	"github.com/nilhiu/srleaderboard/internal/app/middleware"
	"github.com/nilhiu/srleaderboard/internal/view/page"
)

//go:embed static
var staticFS embed.FS

func registerRoutes(ctx context.Context, app *App) {
	authHandler := handler.NewAuthHandler(ctx, app.userSvc)
	validateHandler := handler.NewValidateHandler(ctx)
	runsHandler := handler.NewRunsHandler(ctx, app.userSvc, app.leaderboardSvc)
	userHandler := handler.NewUserPageHandler(ctx)
	partialsHandler := handler.NewPartialsHandler(ctx)
	normalChain := middleware.NewChain(
		middleware.WithJWT(),
		middleware.WithLogging(app.logger),
	)
	protChain := middleware.NewChain(
		middleware.WithProtected(),
		normalChain,
	)

	app.router.Handle("GET /static/", http.FileServerFS(staticFS))

	app.router.Handle("GET /", normalChain.Use(templ.Handler(page.MainPage())))
	app.router.Handle("GET /runs/{user}", normalChain.Use(userHandler.User()))
	app.router.Handle("GET /api/runs", normalChain.Use(runsHandler.GetRuns()))
	app.router.Handle("GET /api/runs/{user}", normalChain.Use(runsHandler.GetUserRuns()))
	app.router.Handle("POST /api/auth/login", normalChain.Use(authHandler.Login()))
	app.router.Handle("POST /api/auth/register", normalChain.Use(authHandler.Register()))
	app.router.Handle("GET /partials/navbar", normalChain.Use(partialsHandler.Navbar()))

	app.router.Handle("GET /profile", protChain.Use(userHandler.Profile()))
	app.router.Handle("POST /api/auth/logout", protChain.Use(authHandler.Logout()))
	app.router.Handle("POST /api/runs", protChain.Use(runsHandler.AddRun()))
	app.router.Handle("POST /partials/validate/time", protChain.Use(validateHandler.Time()))
}
