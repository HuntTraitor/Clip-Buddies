package main

import (
	"log/slog"
	"os"

	"github.com/hunttraitor/clip-buddies/internal/api"
	"github.com/hunttraitor/clip-buddies/internal/auth"
	"github.com/hunttraitor/clip-buddies/internal/config"
	"github.com/hunttraitor/clip-buddies/internal/platform/postgres"
	"github.com/hunttraitor/clip-buddies/internal/system"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := postgres.Open(cfg.DB)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	authHandler := auth.NewHandler(db)
	systemHandler := system.NewHandler()

	app := &api.Application{
		Config:        cfg,
		Logger:        logger,
		DB:            db,
		AuthHandler:   authHandler,
		SystemHandler: systemHandler,
	}

	if err := app.Serve(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
