package core

import (
	"fmt"
	"rest-api/models"
	"rest-api/repositories"
)

// RecipeService is a core service for recipes.
type RecipeService struct {
	DB *repositories.RecipeDb
}

// NewRecipeService creates a new recipe service.
func NewRecipeService(db *repositories.RecipeDb) *RecipeService {
	return &RecipeService{
		DB: db,
	}
}

// SaveRecipe saves a recipe to the database.
func (recipeService *RecipeService) SaveRecipe(recipe *models.Recipe) *models.Response {
	resp := &models.Response{StatusCode: 200}
	if err := recipeService.DB.SaveRecipe(recipe); err != nil {
		resp.StatusCode = 400
		resp.Message = err.Error()
		return resp
	}
	return resp
}

// SaveRecipes saves multiple recipes to the database.
func (recipeService *RecipeService) SaveRecipes(recipes []models.Recipe) *models.Response {
	resp := &models.Response{StatusCode: 200}
	if err := recipeService.DB.SaveRecipes(recipes); err != nil {
		resp.StatusCode = 400
		resp.Message = err.Error()
		return resp
	}
	return resp
}

// DeleteRecipe deletes a recipe by its ID.
func (recipeService *RecipeService) DeleteRecipe(id uint) *models.Response {
	resp := &models.Response{StatusCode: 200}
	if err := recipeService.DB.DeleteRecipe(id); err != nil {
		resp.StatusCode = 400
		resp.Message = err.Error()
		return resp
	}
	return resp
}

// GetRecipe retrieves a recipe by its ID.
func (recipeService *RecipeService) GetRecipe(id uint) *models.Recipe {
	recipe := recipeService.DB.GetRecipe(id)
	if recipe == nil {
		return nil
	}
	return recipe
}

// GetRecipes retrieves all recipes.
func (recipeService *RecipeService) GetRecipes() []*models.Recipe {
	recipes := recipeService.DB.GetRecipes()
	if recipes == nil {
		return nil
	}
	return recipes
}

// GetRecipesByCousines retrieves all recipes by cousines.
func (recipeService *RecipeService) GetRecipesByCousines(cousines []string) []*models.Recipe {
	recipes := recipeService.DB.GetRecipesByCousines(cousines)
	if recipes == nil {
		return nil
	}
	return recipes
}

// GetRecipesByKeywords retrieves all recipes by keywords.
func (recipeService *RecipeService) GetRecipesByKeywords(keywords string) []*models.Recipe {
	recipes := recipeService.DB.GetRecipesByKeywords(keywords)
	if recipes == nil {
		return nil
	}
	return recipes
}

func (recipeService *RecipeService) CompareMarketPrices(recipe *models.Recipe, availableMarkets []models.Market) (*models.Market, error) {
	priceMap := make(map[string]float64)
	uniqueMarkets := trimMarkets(availableMarkets)
	for _, ingredient := range recipe.Ingredients {
		marketIngredients := recipeService.DB.GetMarketIngredientForAllMarkets(ingredient)
		if marketIngredients == nil {
			return nil, fmt.Errorf("error finding market ingredients")
		}
		for _, marketImarketIngredient := range marketIngredients {
			if checkMatchingString(marketImarketIngredient.Market, uniqueMarkets) {
				if val, ok := priceMap[marketImarketIngredient.Market]; ok {
					priceMap[marketImarketIngredient.Market] = val + marketImarketIngredient.Price
				} else {
					priceMap[marketImarketIngredient.Market] = marketImarketIngredient.Price
				}
			}
		}
	}

	minPrice := float64(-1)
	minmarket := ""
	for market, price := range priceMap {
		if minPrice == -1 || price < minPrice {
			minPrice = price
			minmarket = market
		}
	}

	for _, market := range uniqueMarkets {
		if market.Name == minmarket {
			return &market, nil
		}
	}

	return nil, fmt.Errorf("error finding the closest market")
}

func trimMarkets(markets []models.Market) map[string]models.Market {
	// Trim the markets also remove the duplicates
	uniqueMarkets := make(map[string]models.Market)
	for _, market := range markets {
		switch market.Name {
		case "ΣΚΛΑΒΕΝΙΤΗΣ":
			market.Name = "Sklavenitis"
		case "ΑΒ Βασιλόπουλος":
			market.Name = "AB"
		case "Lidl":
			market.Name = "Lidl"
		case "My market":
			market.Name = "MyMarket"
		}

		if _, ok := uniqueMarkets[market.Name]; !ok {
			uniqueMarkets[market.Name] = market
		} else {
			if uniqueMarkets[market.Name].Distance > market.Distance {
				uniqueMarkets[market.Name] = market
			}
		}
	}

	return uniqueMarkets
}

func checkMatchingString(str string, arr map[string]models.Market) bool {
	fmt.Println("Checking for", str, "in", arr)
	for key := range arr {
		if key == str {
			return true
		}
	}
	return false
}

func (recipeService *RecipeService) ConvertRecipeToMarketIngredients(recipe *models.Recipe, market string) []models.MarketIngredient {
	return recipeService.DB.GetMarketIngredientsForMarket(market, recipe)
}

func (s *RecipeService) AddDefaultMarkets() error {
	markets := []string{"Sklavenitis", "AB", "Lidl", "MyMarket"}
	marktesFiles := []string{"Sklavenitis_ingredients.txt", "AB_ingredients.txt", "Lidl_ingredients.txt", "MyMarket_ingredients.txt"}
	for i, market := range markets {
		err := s.DB.AddMarketIngredientsFromFile("SuperMarketPrices/"+marktesFiles[i], market)
		if err != nil {
			return err
		}
	}
	return nil
}

// ΣΚΛΑΒΕΝΙΤΗΣ
// ΑΒ Βασιλόπουλος
// Lidl
// My market
