package localdata

import "backend/models"

type PriceRangeData interface {
	GetAll() []models.PriceRange
}

type priceRangeData struct {
	items []models.PriceRange
}

func InitPriceRangeData() PriceRangeData {
	return &priceRangeData{
		items: []models.PriceRange{
			{ID: 1, DefaultPrice: "0", MinimumPrice: "0", MaximumPrice: "250", Step: "10"},
		},
	}
}

func (p *priceRangeData) GetAll() []models.PriceRange{
	return p.items
}