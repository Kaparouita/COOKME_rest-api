package ports

import (
	"rest-api/models"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	// CreateUser creates a new user.
}

type UserHandler interface {
}

type Db interface {
	NewUserDb()
}

type UserDb interface {
	CreateUser()
	GetUser()
}

type UserService interface {
	CreateUser()
	GetUser()
}

type RecipeDb interface {
	SaveRecipe(*models.Recipe) error
	SaveRecipes([]models.Recipe) error
	GetRecipe(uint) *models.Recipe
	GetRecipes() []*models.Recipe
}

type RecipeService interface {
	SaveRecipe(*models.Recipe) *models.Response
	SaveRecipes([]models.Recipe) *models.Response
	GetRecipe(uint) *models.Recipe
	GetRecipes() []*models.Recipe
}

type RecipeHandler interface {
	SaveRecipe(c *fiber.Ctx)
	SaveRecipes(c *fiber.Ctx)
	GetRecipe(c *fiber.Ctx)
	GetRecipes(c *fiber.Ctx)
}
