package handlers

import (
	"encoding/json"
	"rest-api/models"
	"rest-api/ports"

	"github.com/gofiber/fiber/v2"
)

type RecipeHandler struct {
	srv ports.RecipeService
}

func NewRecipeHandler(srv ports.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		srv: srv,
	}
}

// SaveRecipe saves a recipe to the database.
func (recipeHandler *RecipeHandler) SaveRecipe(c *fiber.Ctx) error {
	recipe := &models.Recipe{}
	err := json.Unmarshal(c.Body(), &recipe)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal Recipe")
	}
	resp := recipeHandler.srv.SaveRecipe(recipe)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON("Saved Successfully")
}

func (recipeHandler *RecipeHandler) SaveRecipes(c *fiber.Ctx) error {
	recipes := []models.Recipe{}
	err := json.Unmarshal(c.Body(), &recipes)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal Recipe")
	}
	resp := recipeHandler.srv.SaveRecipes(recipes)
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON("Saved Successfully")
}
