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
