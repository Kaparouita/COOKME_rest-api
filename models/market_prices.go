package models

type MarketIngredient struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Market string `json:"market"`
}

type Market struct {
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
}
