package server

import (
	"log"
	"rest-api/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	recipeHandler ports.RecipeHandler
	userHandler   ports.UserHandler
	searchHander  ports.SearchHandler
}

func NewService(recipeHandler ports.RecipeHandler, userHandler ports.UserHandler, searchHandler ports.SearchHandler) *Server {
	return &Server{
		recipeHandler: recipeHandler,
		userHandler:   userHandler,
		searchHander:  searchHandler,
	}
}

// Initialize initializes the server by setting up the routes and handlers for various endpoints.
func (server *Server) Initialize() {
	app := fiber.New()
	app.Use(cors.New())

	//recipes
	recipe := app.Group("/recipe")
	recipe.Get("/cuisines", server.recipeHandler.GetRecipesByCousines)
	recipe.Get("/keywords", server.recipeHandler.GetRecipesByKeywords)
	recipe.Post("/convertToMarketIngredients/:recipeID", server.recipeHandler.ConvertRecipeToMarketIngredients)
	recipe.Post("/compareMarketPrices/:recipeID", server.recipeHandler.CompareMarketPrices)
	recipe.Get("/:id", server.recipeHandler.GetRecipe)
	recipe.Delete("/:id", server.recipeHandler.DeleteRecipe)
	recipe.Get("/", server.recipeHandler.GetRecipes)
	recipe.Post("/recipes", server.recipeHandler.SaveRecipes)
	recipe.Post("/", server.recipeHandler.SaveRecipe)

	//search
	search := app.Group("/search")
	search.Get("/:keyword", server.searchHander.Search)
	search.Get("/", server.searchHander.GetAllKeywords)

	//users
	user := app.Group("/user")
	user.Get("/email/:email", server.userHandler.GetUserByEmail)
	user.Get("/all", server.userHandler.GetUsers)
	user.Get("/:id", server.userHandler.GetUser)
	user.Delete("/:id", server.userHandler.DeleteUser)
	user.Post("/", server.userHandler.CreateUser)
	user.Get("/profileRecipes/:userId", server.userHandler.GetProfileRecipes)

	//extra
	user.Get("/closestMarket/:userId", server.userHandler.FindClosestMarket)
	user.Get("/availableMarkets/:userId", server.userHandler.FindAllAvailableMarkets)
	user.Post("/login", server.userHandler.CheckLogin)

	//reviews
	user.Post("/addReview", server.userHandler.AddReview)
	user.Put("/updateReview", server.userHandler.UpdateReview)
	user.Get("/reviews/:userId", server.userHandler.GetReviewsByUserID)

	//favorites
	user.Post("/addFavoriteRecipe/:userId/:recipeId", server.userHandler.AddFavoriteRecipe)
	user.Delete("/removeFavoriteRecipe/:userId/:recipeId", server.userHandler.RemoveFavoriteRecipe)
	user.Get("/favoriteRecipes/:userId", server.userHandler.GetFavoriteRecipes)

	//orders
	order := app.Group("/order")
	order.Get("/all", server.userHandler.GetOrders)
	order.Get("/:userId", server.userHandler.GetOrdersByUserID)
	order.Delete("/:recipeId/:userId", server.userHandler.RemoveOrder)
	order.Post("/", server.userHandler.SaveOrder)

	log.Fatal(app.Listen(":3000"))

}
