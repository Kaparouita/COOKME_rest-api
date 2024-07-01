package models

type Keyword struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Keyword string `json:"keyword"`
}

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
