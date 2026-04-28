package localdata

import "backend/models"

type LogoDropDownData interface {
	GetAll() []models.SectionItem
}

type logoDropDownData struct {
	items []models.SectionItem
}

func InitLogoDropDownData() LogoDropDownData {
	return &logoDropDownData{
		items: []models.SectionItem{
			{
				Items: []models.ItemsData{
					{ID: 1, Label: "👩‍💻 Join Remote OK", Link: "/login"},
					{ID: 2, Label: "👋 Log in", Link: "/login"},
				},
			},
			{
				Title: "General",
				Items: []models.ItemsData{
					{ID: 1, Label: "Frontpage", Link: "/", Image: "/images/logo.webp"},
					{ID: 2, Label: "🏝 Remote jobs", Link: "/"},
					{ID: 3, Label: "🌗 Dark mode", Link: "#"},
					{ID: 4, Label: "👩‍💻 Hire remote workers", Link: "#"},
					{ID: 5, Label: "🚨 Post a job", Link: "#"},
					{ID: 6, Label: "⭐️ Go premium", Link: "#"},
				},
			},
			{
				Title: "Top jobs",
				Items: []models.ItemsData{
					{ID: 1, Label: "🦾 AI Jobs", Link: "#"},
					{ID: 2, Label: "⏰ Async Jobs", Link: "#"},
					{ID: 3, Label: "🌎 Distributed team", Link: "#"},
					{ID: 4, Label: "🎧 Support jobs", Link: "#"},
					{ID: 5, Label: "🤓 Engineer jobs", Link: "#"},
					{ID: 6, Label: "🤓 Software jobs", Link: "#"},
					{ID: 7, Label: "👵 Senior jobs", Link: "#"},
					{ID: 8, Label: "🛠 Technical jobs", Link: "#"},
					{ID: 9, Label: "💼 Management jobs", Link: "#"},
					{ID: 10, Label: "🚀 Growth jobs", Link: "#"},
					{ID: 11, Label: "👩‍✈️ Lead jobs", Link: "#"},
				},
			},
			{
				Title: "Companies",
				Items: []models.ItemsData{
					{ID: 1, Label: "🚨 Post a remote job", Link: "#"},
					{ID: 2, Label: "📦 Buy a job bundle", Link: "#"},
					{ID: 3, Label: "🏷 Ask for a discount", Link: "#"},
					{ID: 4, Label: "Health insurance for teams", Link: "#", Image: "/images/logo-biz-0.png"},
					{ID: 5, Label: "Health insurance for nomads", Link: "#", Image: "/images/logo-biz-0.png"},
				},
			},
			{
				Title: "Feeds",
				Items: []models.ItemsData{
					{ID: 1, Label: "🛠 Remote Jobs API", Link: "#"},
					{ID: 2, Label: "🪚 RSS feed", Link: "#"},
					{ID: 3, Label: "🪓 JSON feed", Link: "#"},
					{ID: 4, Label: "Hacker News mode", Link: "#", Image: "/images/y18.svg"},
					{ID: 5, Label: "Hacker News mode", Link: "#", Image: "/images/vscode.webp"},
				},
			},
			{
				Title: "Help",
				Items: []models.ItemsData{
					{ID: 1, Label: "💡 Ideas + bugs", Link: "#"},
					{ID: 2, Label: "🚀 Changelog", Link: "#"},
					{ID: 3, Label: "🛍️ Merch", Link: "#"},
					{ID: 4, Label: "🛟 FAQ & Help", Link: "#"},
				},
			},
			{
				Title: "Other projects",
				Items: []models.ItemsData{
					{ID: 1, Label: "📊 Remote work stats", Link: "#"},
					{ID: 2, Label: "👷 Top remote companies", Link: "#"},
					{ID: 3, Label: "💰 Highest paying remote jobs", Link: "#"},
					{ID: 4, Label: "🧪 State of remote work", Link: "#"},
					{ID: 5, Label: "🌍 Become a digital nomad", Link: "#"},
					{ID: 6, Label: "🔮 Web3 Jobs", Link: "#"},
					{ID: 7, Label: "📸 Photo AI", Link: "#"},
					{ID: 8, Label: "🏡 Interior AI", Link: "#"},
				},
			},
		},
	}
}

func (l *logoDropDownData) GetAll() []models.SectionItem {
	return l.items
}
