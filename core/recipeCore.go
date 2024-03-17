package core

import (
	"rest-api/models"
	"rest-api/repositories"
)

// RecipeService is a core service for recipes.
type RecipeService struct {
	DB *repositories.RecipeDb
}

// NewRecipeService creates a new recipe service.
func NewRecipeService(db *repositories.RecipeDb) *RecipeService {
	return &RecipeService{
		DB: db,
	}
}

// SaveRecipe saves a recipe to the database.
func (recipeService *RecipeService) SaveRecipe(recipe *models.Recipe) *models.Response {
	resp := &models.Response{StatusCode: 200}
	if err := recipeService.DB.SaveRecipe(recipe); err != nil {
		resp.StatusCode = 400
		resp.Message = err.Error()
		return resp
	}
	return resp
}

// SaveRecipes saves multiple recipes to the database.
func (recipeService *RecipeService) SaveRecipes(recipes []models.Recipe) *models.Response {
	resp := &models.Response{StatusCode: 200}
	if err := recipeService.DB.SaveRecipes(recipes); err != nil {
		resp.StatusCode = 400
		resp.Message = err.Error()
		return resp
	}
	return resp
}

// GetRecipe retrieves a recipe by its ID.
func (recipeService *RecipeService) GetRecipe(id uint) *models.Recipe {
	return recipeService.DB.GetRecipe(id)
}

// GetRecipes retrieves all recipes.
func (recipeService *RecipeService) GetRecipes() []*models.Recipe {
	return recipeService.DB.GetRecipes()
}
