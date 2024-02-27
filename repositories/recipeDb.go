package repositories

import (
	"rest-api/models"

	"github.com/gofiber/fiber/v2"
)

type RecipeDb struct {
	Db *Db
}

func NewRecipeDb(db *Db) *RecipeDb {
	return &RecipeDb{
		Db: db,
	}
}

// SaveRecipe saves a recipe to the database.
func (recipeDb *RecipeDb) SaveRecipe(c *fiber.Ctx) {
	recipe := new(models.Recipe)

	
}
