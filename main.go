package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"rest-api/core"
	"rest-api/models"
	"rest-api/repositories"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := repositories.NewDbRepo()
	// userDb := repositories.NewUserDb(db)
	// userService := core.NewUserService(userDb)
	// userHandler := handlers.NewUserHandler(userService)
	// server := server.NewService(userHandler)

	recipeDb := repositories.NewRecipeDb(db)
	recipeService := core.NewRecipeService(recipeDb)

	// rJson, err := unMarshalRecipes("recipes.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// recipes := models.TranformRecipes(rJson)

	recipes := recipeService.GetRecipes()
	for i, recipe := range recipes {
		fmt.Println(recipe)
		if i == 200 {
			break
		}
	}
	// server.Initialize()
}

func unMarshalRecipes(file string) ([]models.RecipeJson, error) {
	var recipes []models.RecipeJson
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
