package core

import (
	"rest-api/models"
	"rest-api/repositories"
)

type UserService struct {
	DB *repositories.UserDb
}

func NewUserService(db *repositories.UserDb) *UserService {
	return &UserService{
		DB: db,
	}
}

// CreateUser creates a new user.
func (userService *UserService) CreateUser(user *models.User) error {
	return userService.DB.CreateUser(user)
}

// GetUser retrieves a user by their ID.
func (userService *UserService) GetUser(id int) (*models.User, error) {
	return userService.DB.GetUser(id)
}

func (userService *UserService) CheckLogin(login *models.LoginResp) error {
	return userService.DB.CheckLogin(login)
}

func (userService *UserService) GetUserByEmail(email string) (*models.User, error) {
	return userService.DB.GetUserByEmail(email)
}
