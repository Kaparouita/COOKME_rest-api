package core

import "rest-api/repositories"

type UserService struct {
	DB *repositories.UserDb
}

func NewUserService(db *repositories.UserDb) *UserService {
	return &UserService{
		DB: db,
	}
}

// CreateUser creates a new user.
func (userService *UserService) CreateUser() {
	userService.DB.CreateUser()
}

// GetUser retrieves a user by their ID.
func (userService *UserService) GetUser() {
	userService.DB.GetUser()
}
