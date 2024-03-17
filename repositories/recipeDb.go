package repositories

import (
	"encoding/json"
	"rest-api/models"
)

type RecipeDb struct {
	Db *Db
}

func NewRecipeDb(db *Db) *RecipeDb {
	return &RecipeDb{
		Db: db,
	}
}

type RecipeResp struct {
	models.Recipe
	NutritionInfoJson js
}

// SaveRecipe saves a recipe to the database.
func (recipeDb *RecipeDb) SaveRecipe(recipe *models.Recipe) error {
	return recipeDb.Db.Create(recipe).Error
}

func (db *RecipeDb) SaveRecipes(recipes []models.Recipe) error {
	for _, recipe := range recipes {
		err := db.SaveRecipe(&recipe)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *RecipeDb) GetRecipe(id uint) *models.Recipe {
	var recipe models.Recipe
	if err := db.Db.First(&recipe, id).Error; err != nil {
		return nil
	}
	return &recipe
}

func (db *RecipeDb) GetRecipes() []*models.Recipe {
	var recipes []*models.Recipe
	db.Db.Find(&recipes)
	return recipes
}

func unMarshalRecipeNutrition(recipe *models.Recipe) {
	recipe.NutritionInfo = json.Unmarshal(recipe.NutritionInfo)
}
