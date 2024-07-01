package handlers

import (
	"encoding/json"
	"rest-api/models"
	"rest-api/ports"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	srv ports.UserService
}

func NewUserHandler(srv ports.UserService) *UserHandler {
	return &UserHandler{
		srv: srv,
	}
}

// CreateUser creates a new user.
func (userHandler *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := &models.User{}
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal User")
	}
	err = userHandler.srv.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(200).JSON("Saved Successfully")
}

// GetUser retrieves a user by their ID.
func (userHandler *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}
	user, err := userHandler.srv.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("User not found")
	}
	return c.Status(200).JSON(user)
}

func (userHandler *UserHandler) CheckLogin(c *fiber.Ctx) error {
	login := &models.LoginResp{}
	resp := models.Response{}
	err := json.Unmarshal(c.Body(), &login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal Login")
	}
	err = userHandler.srv.CheckLogin(login)
	if err != nil {
		resp.StatusCode = 401
		resp.Message = "Login Failed"
		return c.Status(resp.StatusCode).JSON(resp)
	}
	resp.StatusCode = 200
	resp.Message = "Login Successful"
	return c.Status(resp.StatusCode).JSON(resp)
}

func (userHandler *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := userHandler.srv.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("User not found")
	}
	return c.Status(200).JSON(user)
}
