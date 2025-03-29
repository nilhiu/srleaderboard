package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/nilhiu/srleaderboard/internal/app"
)

func main() {
	ctx := context.Background()
	a, err := app.New(ctx, slog.New(slog.NewTextHandler(os.Stdout, nil)))
	if err != nil {
		panic(err)
	}

	a.Run()
}
