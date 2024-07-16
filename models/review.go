package models

type Review struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	RecipeID int    `json:"recipe_id"`
	UserID   int    `json:"user_id"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
}
