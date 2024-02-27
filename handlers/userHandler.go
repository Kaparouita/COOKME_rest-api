package handlers

import "rest-api/ports"

type UserHandler struct {
	srv ports.UserService
}

func NewUserHandler(srv ports.UserService) *UserHandler {
	return &UserHandler{
		srv: srv,
	}
}

// CreateUser creates a new user.
func (userHandler *UserHandler) CreateUser() {
	userHandler.srv.CreateUser()
}

// GetUser retrieves a user by their ID.
func (userHandler *UserHandler) GetUser() {
	userHandler.srv.GetUser()
}
