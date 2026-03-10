package handlers

import (
	"github.com/hunttraitor/clip-buddies/internal/services"
)

type Handlers struct {
	User   *UserHandler
	System *SystemHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		User:   NewUserHandler(services.User),
		System: NewSystemHandler(services.System),
	}
}
