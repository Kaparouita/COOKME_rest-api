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
	DeleteUser(id int) error
	GetUsers() ([]models.User, error)
	CheckLogin(login *models.LoginResp) error
	GetUserByEmail(email string) (*models.User, error)
	SaveOrder(order *models.Order) error
	GetOrdersByUserID(userID int) ([]models.Order, error)
	AddFavoriteRecipe(userID int, recipeID int64) error
	RemoveFavoriteRecipe(userID int, recipeID int64) error
	GetFavoriteRecipes(userID int) ([]models.Recipe, error)
	AddReview(review *models.Review) error
	GetReviewsByUserID(userID int) ([]models.Review, error)
	RemoveReview(reviewID int) error
	GetProfileRecipes(userID int) ([]models.ProfileRecipeResponse, error)
	UpdateReview(review *models.Review) error
	RemoveOrder(userId, recipeId int) error
	GetOrder(orderID int) (*models.Order, error)
	GetOrders() ([]models.Order, error)
}

type UserService interface {
	CreateUser(*models.User) error
	GetUser(int) (*models.User, error)
	DeleteUser(int) error
	GetUsers() ([]models.User, error)
	CheckLogin(*models.LoginResp) error
	GetUserByEmail(string) (*models.User, error)
	FindAllAvailableMarkets(float64, float64) ([]models.Market, error)
	FindClosestMarket(lan, lon float64) (*models.Market, error)
	SaveOrder(*models.Order) error
	GetOrdersByUserID(int) ([]models.Order, error)
	AddFavoriteRecipe(userID int, recipeID int64) error
	RemoveFavoriteRecipe(userID int, recipeID int64) error
	GetFavoriteRecipes(userID int) ([]models.Recipe, error)
	AddReview(review *models.Review) error
	GetReviewsByUserID(userID int) ([]models.Review, error)
	RemoveReview(reviewID int) error
	GetProfileRecipes(userID int) ([]models.ProfileRecipeResponse, error)
	UpdateReview(review *models.Review) error
	RemoveOrder(userId, recipeId int) error
	GetOrders() ([]models.Order, error)
	GetOrder(orderID int) (*models.Order, error)
}

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	CheckLogin(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error
	FindClosestMarket(c *fiber.Ctx) error
	FindAllAvailableMarkets(c *fiber.Ctx) error
	SaveOrder(c *fiber.Ctx) error
	GetOrdersByUserID(c *fiber.Ctx) error
	AddFavoriteRecipe(c *fiber.Ctx) error
	RemoveFavoriteRecipe(c *fiber.Ctx) error
	GetFavoriteRecipes(c *fiber.Ctx) error
	AddReview(c *fiber.Ctx) error
	GetReviewsByUserID(c *fiber.Ctx) error
	RemoveReview(c *fiber.Ctx) error
	GetProfileRecipes(c *fiber.Ctx) error
	UpdateReview(c *fiber.Ctx) error
	RemoveOrder(c *fiber.Ctx) error
	GetOrders(c *fiber.Ctx) error
}

type RecipeDb interface {
	SaveRecipe(*models.Recipe) error
	DeleteRecipe(uint) error
	SaveRecipes([]models.Recipe) error
	GetRecipe(uint) *models.Recipe
	GetRecipes() []*models.Recipe
	GetRecipesByCousines([]string) []*models.Recipe
	// Keywords are comma separated
	GetRecipesByKeywords(string) []*models.Recipe
	CompareMarketPrices(recipe *models.Recipe, availableMarkets []models.Market) (*models.Market, error)
	GetMarketIngredientsForMarket(market string, recipe *models.Recipe) ([]models.MarketIngredient, error)
}

type RecipeService interface {
	SaveRecipe(*models.Recipe) *models.Response
	DeleteRecipe(uint) *models.Response
	SaveRecipes([]models.Recipe) *models.Response
	GetRecipe(uint) *models.Recipe
	GetRecipes() []*models.Recipe
	GetRecipesByCousines([]string) []*models.Recipe
	// Keywords are comma separated
	GetRecipesByKeywords(string) []*models.Recipe
	CompareMarketPrices(recipe *models.Recipe, availableMarkets []models.Market) (*models.Market, error)
	ConvertRecipeToMarketIngredients(recipe *models.Recipe, market string) []models.MarketIngredient
}

type RecipeHandler interface {
	SaveRecipe(c *fiber.Ctx) error
	DeleteRecipe(c *fiber.Ctx) error
	SaveRecipes(c *fiber.Ctx) error
	GetRecipe(c *fiber.Ctx) error
	GetRecipes(c *fiber.Ctx) error
	GetRecipesByCousines(c *fiber.Ctx) error
	GetRecipesByKeywords(c *fiber.Ctx) error
	CompareMarketPrices(c *fiber.Ctx) error
	ConvertRecipeToMarketIngredients(c *fiber.Ctx) error
}

type SearchHandler interface {
	Search(c *fiber.Ctx) error
	GetAllKeywords(c *fiber.Ctx) error
}

type SearchService interface {
	SearchKeywords(string) ([]string, error)
	GetAllKeywords() []models.Keyword
}
