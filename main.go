package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rest-api/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// db := repositories.NewDbRepo()
	// userDb := repositories.NewUserDb(db)
	// userService := core.NewUserService(userDb)
	// userHandler := handlers.NewUserHandler(userService)

	// server := server.NewService(userHandler)

	rJson, err := unMarshalRecipes("recipes.json")
	if err != nil {
		log.Fatal(err)
	}
	recipes := models.TranformRecipes(rJson)

	uniqueCousines := make(map[string]bool)

	for _, recipe := range recipes {
		uniqueCousines[recipe.Cuisine] = true
	}

	file := "uniqueCousines.txt"

	data := []byte("Unique Cousines\n")
	for k := range uniqueCousines {
		data = append(data, []byte(k+"\n")...)
	}

	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		log.Fatal(err)
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
