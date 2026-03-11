package api

import (
	"database/sql"
	"log/slog"
	"sync"

	"github.com/hunttraitor/clip-buddies/internal/auth"
	"github.com/hunttraitor/clip-buddies/internal/config"
	"github.com/hunttraitor/clip-buddies/internal/system"
)

type Application struct {
	Config config.Config
	Logger *slog.Logger
	DB     *sql.DB
	WG     sync.WaitGroup

	AuthHandler   *auth.Handler
	SystemHandler *system.Handler
}
