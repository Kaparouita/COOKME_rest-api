package main

import (
	"fmt"
	"log"
	"rest-api/core"
	"rest-api/handlers"
	"rest-api/repositories"
	"rest-api/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := repositories.NewDbRepo()

	// User
	userDb := repositories.NewUserDb(db)
	userService := core.NewUserService(userDb)
	userHandler := handlers.NewUserHandler(userService)

	// Recipe
	recipeDb := repositories.NewRecipeDb(db)
	recipeService := core.NewRecipeService(recipeDb)

	market, _ := recipeService.FindClosestMarket(37.979259292519124, 23.771078249581826)
	fmt.Println(market)

	recipeHandler := handlers.NewRecipeHandler(recipeService)

	// Search
	searchService := core.NewSearchService(db)
	searchHandler := handlers.NewSearchHandler(searchService)

	server := server.NewService(recipeHandler, userHandler, searchHandler)

	server.Initialize()
}
