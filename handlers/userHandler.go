package handlers

import (
	"encoding/json"
	"fmt"
	"rest-api/models"
	"rest-api/ports"
	"strconv"

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
	resp := models.Response{}
	user := &models.User{}
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Unable to unmarshal User"
		return c.Status(resp.StatusCode).JSON(resp)
	}
	err = userHandler.srv.CreateUser(user)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Error creating user"
		return c.Status(resp.StatusCode).JSON(resp)
	}

	resp.StatusCode = 200
	resp.Message = "Saved Successfully"
	return c.Status(200).JSON(resp)
}

func (userHandler *UserHandler) DeleteUser(c *fiber.Ctx) error {
	resp := models.Response{}
	id, err := c.ParamsInt("id")
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Invalid ID"
		return c.Status(resp.StatusCode).JSON(resp)
	}
	err = userHandler.srv.DeleteUser(id)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Error deleting user"
		return c.Status(resp.StatusCode).JSON(resp)
	}
	resp.StatusCode = 200
	resp.Message = "Deleted Successfully"
	return c.Status(200).JSON(resp)
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

func (userHandler *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := userHandler.srv.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding users")
	}
	return c.Status(200).JSON(users)
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

func (userHandler *UserHandler) FindClosestMarket(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	user, err := userHandler.srv.GetUser(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("User not found")
	}

	market, err := userHandler.srv.FindClosestMarket(user.Latitude, user.Longitude)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding closest market")
	}

	return c.Status(200).JSON(market)
}

func (userHandler *UserHandler) FindAllAvailableMarkets(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	user, err := userHandler.srv.GetUser(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("User not found")
	}

	markets, err := userHandler.srv.FindAllAvailableMarkets(user.Latitude, user.Longitude)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding available markets")
	}

	return c.Status(200).JSON(markets)
}

func (userHandler *UserHandler) SaveOrder(c *fiber.Ctx) error {
	order := &models.Order{}
	err := json.Unmarshal(c.Body(), &order)
	fmt.Println(order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal Order")
	}
	err = userHandler.srv.SaveOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	resp := models.Response{
		StatusCode: 200,
		Message:    "Saved Successfully",
	}
	return c.Status(200).JSON(resp)
}

func (userHandler *UserHandler) GetOrders(c *fiber.Ctx) error {
	orders, err := userHandler.srv.GetOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding orders")
	}

	return c.Status(200).JSON(orders)
}

func (userHandler *UserHandler) GetOrdersByUserID(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	orders, err := userHandler.srv.GetOrdersByUserID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding orders")
	}

	return c.Status(200).JSON(orders)
}

func (userHandler *UserHandler) AddReview(c *fiber.Ctx) error {
	review := &models.Review{}
	err := json.Unmarshal(c.Body(), &review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal Review")
	}
	err = userHandler.srv.AddReview(review)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	resp := models.Response{
		StatusCode: 200,
		Message:    "Saved Successfully",
	}
	return c.Status(200).JSON(resp)
}

func (userHandler *UserHandler) GetReviewsByUserID(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	reviews, err := userHandler.srv.GetReviewsByUserID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding reviews")
	}

	return c.Status(200).JSON(reviews)
}

func (userHandler *UserHandler) UpdateReview(c *fiber.Ctx) error {
	review := &models.Review{}
	err := json.Unmarshal(c.Body(), &review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal Review")
	}
	err = userHandler.srv.UpdateReview(review)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	resp := models.Response{
		StatusCode: 201,
		Message:    "Saved Successfully",
	}
	return c.Status(201).JSON(resp)
}

func (userHandler *UserHandler) AddFavoriteRecipe(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	recipeId, err := strconv.ParseInt(c.Params("recipeId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	err = userHandler.srv.AddFavoriteRecipe(userId, recipeId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	resp := models.Response{
		StatusCode: 200,
		Message:    "Saved Successfully",
	}

	return c.Status(200).JSON(resp)
}

func (userHandler *UserHandler) RemoveFavoriteRecipe(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	recipeId, err := strconv.ParseInt(c.Params("recipeId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	err = userHandler.srv.RemoveFavoriteRecipe(userId, recipeId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	resp := models.Response{
		StatusCode: 200,
		Message:    "Saved Successfully",
	}

	return c.Status(200).JSON(resp)
}

func (userHandler *UserHandler) RemoveReview(c *fiber.Ctx) error {
	reviewId, err := strconv.Atoi(c.Params("reviewId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	err = userHandler.srv.RemoveReview(reviewId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	resp := models.Response{
		StatusCode: 200,
		Message:    "Saved Successfully",
	}

	return c.Status(200).JSON(resp)
}

func (userHandler *UserHandler) GetFavoriteRecipes(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	recipes, err := userHandler.srv.GetFavoriteRecipes(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding favorite recipes")
	}

	return c.Status(200).JSON(recipes)
}

func (userHandler *UserHandler) GetProfileRecipes(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	recipes, err := userHandler.srv.GetProfileRecipes(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error finding profile recipes")
	}

	return c.Status(200).JSON(recipes)
}

func (userHandler *UserHandler) RemoveOrder(c *fiber.Ctx) error {
	recipeId, err := strconv.Atoi(c.Params("recipeId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}
	userId, err := strconv.Atoi(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	err = userHandler.srv.RemoveOrder(userId, recipeId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	resp := models.Response{
		StatusCode: 200,
		Message:    "Removed Successfully",
	}

	return c.Status(200).JSON(resp)
}
