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
	recipe.Get("/:id", server.recipeHandler.GetRecipe)
	recipe.Get("/", server.recipeHandler.GetRecipes)

	//search
	search := app.Group("/search")
	search.Get("/:keyword", server.searchHander.Search)
	search.Get("/", server.searchHander.GetAllKeywords)

	//users
	user := app.Group("/user")
	user.Get("/email/:email", server.userHandler.GetUserByEmail) // Define the email route first to avoid conflicts
	user.Post("/login", server.userHandler.CheckLogin)
	user.Get("/:id", server.userHandler.GetUser)
	user.Post("/", server.userHandler.CreateUser)

	log.Fatal(app.Listen(":3000"))

}
