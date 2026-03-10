package handlers

import "github.com/hunttraitor/clip-buddies/internal/user"

type UserHandler struct {
	UserService *user.Service
}

func NewUserHandler(userService *user.Service) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
