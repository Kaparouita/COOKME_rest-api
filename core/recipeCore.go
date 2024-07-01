package core

import (
	"context"
	"fmt"
	"rest-api/models"
	"rest-api/repositories"

	"googlemaps.github.io/maps"
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

func trimMarkets(markets []maps.PlacesSearchResult) []maps.PlacesSearchResult {
	trimmedMarkets := make([]maps.PlacesSearchResult, 0)
	for _, market := range markets {
		if market.Name == "ΣΚΛΑΒΕΝΙΤΗΣ" || market.Name == "ΑΒ Βασιλόπουλος" || market.Name == "Lidl" || market.Name == "My market" {
			trimmedMarkets = append(trimmedMarkets, market)
		}
	}
	return trimmedMarkets
}

func (recipeService *RecipeService) SpecificClosestMarket(lan, lon float64, marketName string) (*models.Market, error) {
	// Your Google Maps
	apiKey := "AIzaSyCULKAWrNKbSkXtFx74rn80_sZpOlc4R7U"

	// Initialize the client
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating the client: %s", err)
	}

	// Given location (latitude, longitude)
	givenLocation := maps.LatLng{
		Lat: lan,
		Lng: lon,
	}

	// Find nearby markets
	markets, err := findNearbyMarkets(client, givenLocation)
	if err != nil {
		return nil, fmt.Errorf("error finding nearby markets: %s", err)
	}

	// Trim the markets
	markets = trimMarkets(markets)
	specificMarkets := make([]maps.PlacesSearchResult, 0)

	for _, market := range markets {
		if market.Name == marketName {
			specificMarkets = append(specificMarkets, market)
		}
	}

	// Find the closest market
	closestMarket, distance, err := findClosestMarket(client, givenLocation, specificMarkets)
	if err != nil {
		return nil, fmt.Errorf("error finding the closest market: %s", err)
	}

	return &models.Market{
		Name:     closestMarket.Name,
		Distance: distance,
	}, nil
}

func (recipeService *RecipeService) CompareMarketPrices(recipe *models.Recipe) (*models.Market, error) {
	marketIngredients := make(map[string]float64)
	for _, ingredient := range recipe.Ingredients {
		marketIngredient := recipeService.DB.GetMarketIngredientForAllMarkets(ingredient)
		if marketIngredient == nil {
			return nil, fmt.Errorf("error finding market ingredients")
		}
		for _, market := range marketIngredient {
			if val, ok := marketIngredients[market.Name]; ok {
				marketIngredients[market.Name] = val + market.Price
			} else {
				marketIngredients[market.Name] = market.Price
			}
		}
	}

	return &models.Market{
		Name:     "All Markets",
		Prices:   marketIngredients,
		Distance: 0,
	}, nil
	
}

func (recipeService *RecipeService) AddMarketIngredientsFromFile(path, market string) error {

func (recipeService *RecipeService) FindClosestMarket(lan, lon float64) (*models.Market, error) {
	// Your Google Maps API key
	apiKey := "AIzaSyCULKAWrNKbSkXtFx74rn80_sZpOlc4R7U"

	// Initialize the client
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating the client: %s", err)
	}

	// Given location (latitude, longitude)
	givenLocation := maps.LatLng{
		Lat: lan,
		Lng: lon,
	}

	// Find nearby markets
	markets, err := findNearbyMarkets(client, givenLocation)
	if err != nil {
		return nil, fmt.Errorf("error finding nearby markets: %s", err)
	}

	// Trim the markets
	markets = trimMarkets(markets)

	// Find the closest market
	closestMarket, distance, err := findClosestMarket(client, givenLocation, markets)
	if err != nil {
		return nil, fmt.Errorf("error finding the closest market: %s", err)
	}

	return &models.Market{
		Name:     closestMarket.Name,
		Distance: distance,
	}, nil
}

func findNearbyMarkets(client *maps.Client, location maps.LatLng) ([]maps.PlacesSearchResult, error) {
	// Search for nearby supermarkets 5 km around the given location and 20 results
	r := &maps.NearbySearchRequest{
		Location: &location,
		Radius:   5000,          // Search within 5 km radius
		Type:     "supermarket", // Type of place to search for
		OpenNow:  true,          // Only search for places that are open no

	}

	resp, err := client.NearbySearch(context.Background(), r)
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}

func findClosestMarket(client *maps.Client, origin maps.LatLng, markets []maps.PlacesSearchResult) (maps.PlacesSearchResult, float64, error) {
	var closestMarket maps.PlacesSearchResult
	minDistance := float64(-1)

	for _, market := range markets {
		destination := maps.LatLng{Lat: market.Geometry.Location.Lat, Lng: market.Geometry.Location.Lng}

		r := &maps.DistanceMatrixRequest{
			Origins:      []string{fmt.Sprintf("%f,%f", origin.Lat, origin.Lng)},
			Destinations: []string{fmt.Sprintf("%f,%f", destination.Lat, destination.Lng)},
		}

		resp, err := client.DistanceMatrix(context.Background(), r)
		if err != nil {
			return maps.PlacesSearchResult{}, 0, err
		}

		// Get the distance in meters
		distance := resp.Rows[0].Elements[0].Distance.Meters

		fmt.Printf("Distance to %s: %d meters\n", market.Name, distance)

		if minDistance == -1 || float64(distance) < minDistance {
			minDistance = float64(distance)
			closestMarket = market
		}
	}

	return closestMarket, minDistance, nil
}

// ΣΚΛΑΒΕΝΙΤΗΣ
// ΑΒ Βασιλόπουλος
// Lidl
// My market
