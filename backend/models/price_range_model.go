package models

type PriceRange struct {
	ID int `json:"id"`
	DefaultPrice string `json:"defaultPrice"`
	MinimumPrice string `json:"minimumPrice"`
	MaximumPrice string `json:"maximumPrice"`
	Step string `json:"step"`
}