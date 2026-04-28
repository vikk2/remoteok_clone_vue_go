package localdata

import "backend/models"

type SortByData interface {
	GetAll() []models.SortByData
}

type sortByData struct {
	items []models.SortByData
}

func InitSortByData() SortByData {
	return &sortByData{
		items: []models.SortByData{
			{ID: 1, Value: "default", Item: "🦴 Sort by"},
			{ID: 2, Value: "new", Item: "🆕 Latest jobs"},
			{ID: 3, Value: "high", Item: "💵 Highest paid"},
			{ID: 4, Value: "most", Item: "👀 Most viewed"},
			{ID: 5, Value: "most-applied", Item: "✅ Most applied"},
			{ID: 6, Value: "hot", Item: "🔥 Hottest"},
			{ID: 7, Value: "benefits", Item: "🎪 Most benefits"},
		},
	}
}

func (s *sortByData) GetAll() []models.SortByData {
	return s.items
}
