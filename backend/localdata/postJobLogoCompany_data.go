package localdata

import "backend/models"

type PostJobLogoData interface {
	GetAll() []models.PostJobLogoData
}

type postJobLogoData struct {
	items []models.PostJobLogoData
}

func InitPostJobLogoData() PostJobLogoData{
	return &postJobLogoData{
		items: []models.PostJobLogoData{
			{ID: 1, Image: "/images/aws.webp"},
			{ID: 2, Image: "/images/espn.svg"},
			{ID: 3, Image: "/images/ibm.webp"},
			{ID: 4, Image: "/images/microsoft.webp"},
			{ID: 5, Image: "/images/scale-ai.webp"},
			{ID: 6, Image: "/images/ycombinator.png"},
			{ID: 7, Image: "/images/github.webp"},
			{ID: 8, Image: "/images/posthog.webp"},
			{ID: 9, Image: "/images/mozilla.png"},
			{ID: 10, Image: "/images/stripe.png"},
			{ID: 11, Image: "/images/cloudflare.webp"},
			{ID: 12, Image: "/images/shopify.webp"},
			{ID: 13, Image: "/images/godaddy.webp"},
			{ID: 14, Image: "/images/easyjet.webp"},
			{ID: 15, Image: "/images/indeed.webp"},
		},
	}
}

func (p *postJobLogoData) GetAll() []models.PostJobLogoData{
	return p.items
}