package ports

import (
	"rest-api/models"

	"github.com/gofiber/fiber/v2"
)

type Db interface {
	NewUserDb()
	GetAllKeywords() []models.Keyword
}

type UserDb interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
	CheckLogin(login *models.LoginResp) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserService interface {
	CreateUser(*models.User) error
	GetUser(int) (*models.User, error)
	CheckLogin(*models.LoginResp) error
	GetUserByEmail(string) (*models.User, error)
}

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	CheckLogin(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error
}

type RecipeDb interface {
	SaveRecipe(*models.Recipe) error
	SaveRecipes([]models.Recipe) error
	GetRecipe(uint) *models.Recipe
	GetRecipes() []*models.Recipe
	GetRecipesByCousines([]string) []*models.Recipe
	// Keywords are comma separated
	GetRecipesByKeywords(string) []*models.Recipe
}

type RecipeService interface {
	SaveRecipe(*models.Recipe) *models.Response
	SaveRecipes([]models.Recipe) *models.Response
	GetRecipe(uint) *models.Recipe
	GetRecipes() []*models.Recipe
	GetRecipesByCousines([]string) []*models.Recipe
	// Keywords are comma separated
	GetRecipesByKeywords(string) []*models.Recipe
	FindClosestMarket(lan, lon float64) (models.Market, error)
}

type RecipeHandler interface {
	SaveRecipe(c *fiber.Ctx) error
	SaveRecipes(c *fiber.Ctx) error
	GetRecipe(c *fiber.Ctx) error
	GetRecipes(c *fiber.Ctx) error
	GetRecipesByCousines(c *fiber.Ctx) error
	GetRecipesByKeywords(c *fiber.Ctx) error
}

type SearchHandler interface {
	Search(c *fiber.Ctx) error
	GetAllKeywords(c *fiber.Ctx) error
}

type SearchService interface {
	SearchKeywords(string) ([]string, error)
	GetAllKeywords() []models.Keyword
}
