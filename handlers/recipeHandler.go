package handlers

import (
	"encoding/json"
	"rest-api/models"
	"rest-api/ports"
	"strconv"
	"strings"

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

	return c.Status(resp.StatusCode).JSON(resp)
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

	return c.Status(resp.StatusCode).JSON(resp)
}

func (recipeHandler *RecipeHandler) DeleteRecipe(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}
	resp := recipeHandler.srv.DeleteRecipe(uint(id))
	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp)
}

// GetRecipes returns all recipes from the database.
func (recipeHandler *RecipeHandler) GetRecipes(c *fiber.Ctx) error {
	recipes := recipeHandler.srv.GetRecipes()
	if recipes == nil {
		return c.Status(404).JSON("No recipes found")
	}
	return c.Status(200).JSON(recipes)
}

// GetRecipe returns a recipe by its ID.
func (recipeHandler *RecipeHandler) GetRecipe(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}
	recipe := recipeHandler.srv.GetRecipe(uint(id))
	if recipe == nil {
		return c.Status(404).JSON("Recipe not found")
	}
	return c.Status(200).JSON(recipe)
}

func (recipeHandler *RecipeHandler) GetRecipesByCousines(c *fiber.Ctx) error {
	// get the cousines as an array of strings
	cuisinesParam := c.Query("cuisines")
	if cuisinesParam == "" {
		return c.Status(400).JSON("Invalid cousines")
	}
	cuisines := strings.Split(cuisinesParam, ",")
	recipes := recipeHandler.srv.GetRecipesByCousines(cuisines)
	if recipes == nil {
		return c.Status(404).JSON("No recipes found")
	}
	return c.Status(200).JSON(recipes)
}

func (recipeHandler *RecipeHandler) GetRecipesByKeywords(c *fiber.Ctx) error {
	keywordsParam := c.Query("keywords")
	if keywordsParam == "" {
		return c.Status(400).JSON("Invalid keywords")
	}
	recipes := recipeHandler.srv.GetRecipesByKeywords(keywordsParam)
	if recipes == nil {
		return c.Status(404).JSON("No recipes found")
	}
	return c.Status(200).JSON(recipes)
}

func (recipeHandler *RecipeHandler) CompareMarketPrices(c *fiber.Ctx) error {
	recipeID, err := strconv.Atoi(c.Params("recipeID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	recipe := recipeHandler.srv.GetRecipe(uint(recipeID))
	if recipe == nil {
		return c.Status(fiber.StatusNotFound).JSON("Recipe not found")
	}

	// unmashal the markets
	markets := []models.Market{}
	err = json.Unmarshal(c.Body(), &markets)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to unmarshal Markets")
	}

	market, err := recipeHandler.srv.CompareMarketPrices(recipe, markets)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Error comparing market prices")
	}

	return c.Status(200).JSON(market)
}

func (recipeHandler *RecipeHandler) ConvertRecipeToMarketIngredients(c *fiber.Ctx) error {
	recipeID, err := strconv.Atoi(c.Params("recipeID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid ID")
	}

	market := c.Query("market")
	if market == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid market")
	}

	recipe := recipeHandler.srv.GetRecipe(uint(recipeID))
	if recipe == nil {
		return c.Status(fiber.StatusNotFound).JSON("Recipe not found")
	}

	marketIngredients := recipeHandler.srv.ConvertRecipeToMarketIngredients(recipe, market)

	return c.Status(200).JSON(marketIngredients)
}
