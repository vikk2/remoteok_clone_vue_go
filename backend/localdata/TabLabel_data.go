package localdata

import "backend/models"

type LabelData interface {
	GetAll() []models.TabLabelData
}

type labelData struct {
	items []models.TabLabelData
}

func InitLabelData() LabelData {
	return &labelData{
		items: []models.TabLabelData{
			{
				"🎧 Support",
				"🤓 Engineer",
				"🤓 Software",
				"👵 Senior",
				"🛠 Technical",
				"💼 Management",
				"🚀 Growth",
				"👩‍✈️ Lead",
			},
		},
	}
}

func (l *labelData) GetAll() []models.TabLabelData {
	return l.items
}
