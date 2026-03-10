package services

import (
	"database/sql"

	"github.com/hunttraitor/clip-buddies/internal/system"
	"github.com/hunttraitor/clip-buddies/internal/user"
)

type Services struct {
	User   *user.Service
	System *system.Service
}

func NewServices(db *sql.DB) *Services {
	return &Services{
		User:   user.New(db),
		System: system.New(db),
	}
}
