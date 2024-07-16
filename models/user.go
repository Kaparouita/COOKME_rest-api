package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	Id              uint          `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time     `json:"created_at"`
	FirstName       string        `json:"first_name"`
	LastName        string        `json:"last_name"`
	Username        string        `json:"username" gorm:"uniqueIndex"` // Username (un)
	Password        string        `json:"password"`                    // Password (confirm)
	Email           string        `json:"email" gorm:"uniqueIndex"`    // Email (un)
	Gender          string        `json:"gender"`
	UserType        string        `json:"user_type"` // Type (admin/user/IngridientAdmin)
	Address         string        `json:"address"`
	Phone           string        `json:"phone"`
	URL             string        `json:"url"` // URL (fb / insta)
	Longitude       float64       `json:"longitude"`
	Latitude        float64       `json:"latitude"`
	FavoriteRecipes pq.Int64Array `gorm:"type:integer[]" json:"favorites_recipes"`
	Response
}

type LoginResp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
