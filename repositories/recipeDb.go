package repositories

import (
	"bufio"
	"fmt"
	"os"
	"rest-api/models"
	"strconv"
	"strings"
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
	// NutritionInfoJson js
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
	recipe.UnmarshalNutritionInfo()
	return &recipe
}

func (db *RecipeDb) DeleteRecipe(id uint) error {
	return db.Db.Delete(&models.Recipe{}, id).Error
}

func (db *RecipeDb) GetRecipes() []*models.Recipe {
	var recipes []*models.Recipe
	err := db.Db.Find(&recipes).Error
	if err != nil {
		return nil
	}

	for _, recipe := range recipes {
		recipe.UnmarshalNutritionInfo()
	}
	return recipes
}

func (db *RecipeDb) GetRecipesByCousines(cuisines []string) []*models.Recipe {
	var recipes []*models.Recipe
	fmt.Println("Retrieving recipes by cuisines : ", cuisines)
	err := db.Db.Where("cuisine IN (?)", cuisines).Find(&recipes).Error
	if err != nil {
		return nil
	}

	for _, recipe := range recipes {
		recipe.UnmarshalNutritionInfo()
	}
	return recipes
}

func (db *RecipeDb) GetRecipesByKeywords(keywords string) []*models.Recipe {
	var recipes []*models.Recipe
	fmt.Println("Retrieving recipes by keywords:", keywords)

	keywordArray := "{" + keywords + "}"

	// Execute the query
	err := db.Db.Where("keywords && ?", keywordArray).Find(&recipes).Error
	if err != nil {
		fmt.Println("Error retrieving recipes:", err)
		return nil
	}

	for _, recipe := range recipes {
		recipe.UnmarshalNutritionInfo()
	}
	return recipes
}

func (db *RecipeDb) GetAllKeywords() []models.Keyword {
	var keywords []models.Keyword
	err := db.Db.Find(&keywords).Error
	if err != nil {
		return nil
	}
	return keywords
}

func (db *RecipeDb) GetMarketIngredientForAllMarkets(ingredient string) []models.MarketIngredient {
	var marketIngredients []models.MarketIngredient
	err := db.Db.Where("name = ?", ingredient).Find(&marketIngredients).Error
	if err != nil {
		return nil
	}

	return marketIngredients
}

func (db *RecipeDb) AddMarketIngredientsFromFile(path, market string) error {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// Read the file line by line Ingredient,Price
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into ingredient and price
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 2 {
			continue
		}
		price, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return err
		}
		ingredient := models.MarketIngredient{
			Name:   parts[0],
			Price:  price,
			Market: market,
		}
		// Insert the ingredient
		db.Db.Create(&ingredient)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (db *RecipeDb) GetMarketIngredientsForMarket(market string, recipe *models.Recipe) []models.MarketIngredient {
	var marketIngredients []models.MarketIngredient
	for _, ingredient := range recipe.Ingredients {
		var marketIngredient models.MarketIngredient
		db.Db.Where("name = ? AND market = ?", ingredient, market).First(&marketIngredient)
		marketIngredients = append(marketIngredients, marketIngredient)
	}
	return marketIngredients
}
