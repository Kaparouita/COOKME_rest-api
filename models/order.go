package models

type Order struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	UserID   int     `json:"user_id"`
	RecipeID int     `json:"recipe_id"`
	Market   string  `json:"market"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
	OrderNow bool    `json:"order_now"`
	Date     string  `json:"date"`
}
