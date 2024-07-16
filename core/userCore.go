package core

import (
	"context"
	"fmt"
	"rest-api/models"
	"rest-api/repositories"

	"googlemaps.github.io/maps"
)

type UserService struct {
	DB *repositories.UserDb
}

func NewUserService(db *repositories.UserDb) *UserService {
	return &UserService{
		DB: db,
	}
}

// CreateUser creates a new user.
func (userService *UserService) CreateUser(user *models.User) error {
	return userService.DB.CreateUser(user)
}

// GetUser retrieves a user by their ID.
func (userService *UserService) GetUser(id int) (*models.User, error) {
	return userService.DB.GetUser(id)
}

// DeleteUser deletes a user by their ID.
func (userService *UserService) DeleteUser(id int) error {
	return userService.DB.DeleteUser(id)
}

func (userService *UserService) GetUsers() ([]models.User, error) {
	return userService.DB.GetUsers()
}

func (userService *UserService) CheckLogin(login *models.LoginResp) error {
	return userService.DB.CheckLogin(login)
}

func (userService *UserService) GetUserByEmail(email string) (*models.User, error) {
	return userService.DB.GetUserByEmail(email)
}

func (s *UserService) FindAllAvailableMarkets(lat, lon float64) ([]models.Market, error) {
	// Your Google Maps API key
	apiKey := "AIzaSyCULKAWrNKbSkXtFx74rn80_sZpOlc4R7U"

	// Initialize the client
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating the client: %s", err)
	}

	// Given location (latitude, longitude)
	givenLocation := maps.LatLng{
		Lat: lat,
		Lng: lon,
	}

	// Find nearby markets
	markets, err := findNearbyMarkets(client, givenLocation)
	if err != nil {
		return nil, fmt.Errorf("error finding nearby markets: %s", err)
	}

	// Trim the markets
	markets = chooseMarkets(markets)

	fmt.Println("Found the following markets:")
	for _, market := range markets {
		fmt.Println(market.Name)
	}

	// Convert the markets to the response format
	var response []models.Market
	for _, market := range markets {
		// Also calculate the distance
		destination := maps.LatLng{Lat: market.Geometry.Location.Lat, Lng: market.Geometry.Location.Lng}

		r := &maps.DistanceMatrixRequest{
			Origins:      []string{fmt.Sprintf("%f,%f", givenLocation.Lat, givenLocation.Lng)},
			Destinations: []string{fmt.Sprintf("%f,%f", destination.Lat, destination.Lng)},
		}

		resp, err := client.DistanceMatrix(context.Background(), r)
		if err != nil {
			return nil, fmt.Errorf("error calculating the distance: %s", err)
		}

		// Get the distance in meters
		distance := resp.Rows[0].Elements[0].Distance.Meters

		response = append(response, models.Market{
			Name:     market.Name,
			Distance: float64(distance),
		})
	}

	return response, nil
}

func (s *UserService) FindClosestMarket(lan, lon float64) (*models.Market, error) {
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
	markets = chooseMarkets(markets)

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

	}

	resp, err := client.NearbySearch(context.Background(), r)
	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}

func chooseMarkets(markets []maps.PlacesSearchResult) []maps.PlacesSearchResult {
	trimmedMarkets := make([]maps.PlacesSearchResult, 0)
	for _, market := range markets {
		if market.Name == "ΣΚΛΑΒΕΝΙΤΗΣ" || market.Name == "ΑΒ Βασιλόπουλος" || market.Name == "Lidl" || market.Name == "My market" {
			market.Name = changeNameToBasic(market.Name)
			trimmedMarkets = append(trimmedMarkets, market)
		}
	}
	return trimmedMarkets
}

func changeNameToBasic(name string) string {
	switch name {
	case "ΣΚΛΑΒΕΝΙΤΗΣ":
		return "Sklavenitis"
	case "ΑΒ Βασιλόπουλος":
		return "AB"
	case "Lidl":
		return "Lidl"
	case "My market":
		return "MyMarket"
	}
	return name
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

func (s *UserService) SaveOrder(order *models.Order) error {
	return s.DB.SaveOrder(order)
}

func (s *UserService) GetOrdersByUserID(userID int) ([]models.Order, error) {
	return s.DB.GetOrdersByUserID(userID)
}

func (s *UserService) AddFavoriteRecipe(userID int, recipeID int64) error {
	return s.DB.AddFavoriteRecipe(userID, recipeID)
}

func (s *UserService) GetFavoriteRecipes(userID int) ([]models.Recipe, error) {
	return s.DB.GetFavoriteRecipes(userID)
}

func (s *UserService) RemoveFavoriteRecipe(userID int, recipeID int64) error {
	return s.DB.RemoveFavoriteRecipe(userID, recipeID)
}

func (s *UserService) AddReview(review *models.Review) error {
	return s.DB.AddReview(review)
}

func (s *UserService) GetReviewsByUserID(userID int) ([]models.Review, error) {
	return s.DB.GetReviewsByUserID(userID)
}

func (s *UserService) UpdateReview(review *models.Review) error {
	return s.DB.UpdateReview(review)
}

func (s *UserService) RemoveReview(reviewID int) error {
	return s.DB.RemoveReview(reviewID)
}

func (s *UserService) GetProfileRecipes(userID int) ([]models.ProfileRecipeResponse, error) {
	recipes, err := s.DB.GetRecipes(userID)
	if err != nil {
		return nil, err
	}
	reviews, err := s.DB.GetReviewsByUserID(userID)
	if err != nil {
		return nil, err
	}
	favorites, err := s.DB.GetFavoriteRecipes(userID)
	if err != nil {
		return nil, err
	}

	profileResp := make([]models.ProfileRecipeResponse, 0)
	for _, recipe := range recipes {
		review := getReviewsForRecipe(int(recipe.Id), reviews)
		favorite := isFavorite(int(recipe.Id), favorites)
		profileResp = append(profileResp, models.ProfileRecipeResponse{
			RecipeID:    int(recipe.Id),
			RecipeName:  recipe.Title,
			RecipeImage: recipe.Image,
			Review:      review,
			TotalPrice:  recipe.Price,
			OrderStatus: recipe.OrderStatus,
			IsFavorite:  favorite,
		})
	}

	return profileResp, nil
}

func getReviewsForRecipe(recipeID int, reviews []models.Review) models.Review {
	var res models.Review
	for _, review := range reviews {
		if int(review.RecipeID) == recipeID {
			res = review
			break
		}
	}
	return res
}

func isFavorite(recipeID int, favorites []models.Recipe) bool {
	for _, favorite := range favorites {
		if int(favorite.Id) == recipeID {
			return true
		}
	}
	return false
}

func (s *UserService) RemoveOrder(userId, recipeId int) error {
	return s.DB.RemoveOrder(userId, recipeId)
}

func (s *UserService) GetOrders() ([]models.Order, error) {
	return s.DB.GetOrders()
}

func (s *UserService) GetOrder(orderID int) (*models.Order, error) {
	return s.DB.GetOrder(orderID)
}
