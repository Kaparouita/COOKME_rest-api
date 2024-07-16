package models

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type RecipeResponse struct {
	StatusCode        int                `json:"status_code"`
	Message           string             `json:"message"`
	MarketIngredients []MarketIngredient `json:"market_ingredients"`
	Market            Market             `json:"market"`
}

type ProfileRecipeResponse struct {
	RecipeID    int     `json:"recipe_id"`
	RecipeName  string  `json:"recipe_name"`
	RecipeImage string  `json:"recipe_image"`
	TotalPrice  float64 `json:"total_price"`
	OrderStatus string  `json:"order_status"`
	Review      Review  `json:"review"`
	IsFavorite  bool    `json:"is_favorite"`
}
