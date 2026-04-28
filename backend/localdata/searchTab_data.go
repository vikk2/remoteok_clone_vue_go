package localdata

import "backend/models"

type SearchTabData interface {
	GetAll() []models.SearchTabData
}

type searchTabData struct {
	items []models.SearchTabData
}

func InitSearchTabData() SearchTabData{
	return &searchTabData{
		items: []models.SearchTabData{
			{
				"🎧 Support",
				"🤓 Engineer",
				"🤓 Software",
				"👵 Senior",
				"🛠 Technical",
				"💼 Management",
				"🚀 Growth",
				"👩‍✈️ Lead",
				"🎨 Design",
				"💼 Manager",
				"💼 Sales",
				"💰 Financial",
				"🚥 Marketing",
				"🔑 Security",
				"♾️ Operations",
				"☁️ Cloud",
				"♾️ System",
				"♾️ Operational",
				"💼 Strategy",
				"💼 Leader",
				"🤓 Code",
				"🔌 Non Tech",
				"👨‍🏫 Training",
				"✍️ Content",
			},
		},
	}
}

func (s *searchTabData) GetAll() []models.SearchTabData{
	return s.items
}