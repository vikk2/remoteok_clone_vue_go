package localdata

import "backend/models"

type RightSideInfoData interface {
	GetAll() []models.RightSideInfo
}

type rightSideInfoData struct {
	items []models.RightSideInfo
}

func InitRightSideInfoData() RightSideInfoData {
	return &rightSideInfoData{
		items: []models.RightSideInfo{
			{
				Label:       "Starting from",
				ValueText:   "💳 $299",
				ValueImages: []models.RightSideValueImage{},
				ValueClass:  "",
				Suffix:      "for 30 days",
				BoxClass:    "",
				LinkPrefix:  "",
				LinkText:    "",
				LinkHref:    "#",
			},
			{
				Label:       "Reach",
				ValueText:   "🚀 2,600,000+",
				ValueImages: []models.RightSideValueImage{},
				ValueClass:  "",
				Suffix:      "remote workers/mo",
				BoxClass:    "",
				LinkPrefix:  "",
				LinkText:    "",
				LinkHref:    "#",
			},
			{
				Label:     "Distributed on the",
				ValueText: "Google for Jobs",
				ValueImages: []models.RightSideValueImage{
					{Src: "/images/google-svg.svg", Alt: "google-svg", Width: 15},
				},
				ValueClass: "",
				Suffix:     "recruitment network",
				BoxClass:   "",
				LinkPrefix: "",
				LinkText:   "",
				LinkHref:   "#",
			},
			{
				Label:     "Rated",
				ValueText: "",
				ValueImages: []models.RightSideValueImage{
					{Src: "/images/emoji-star.webp", Alt: "star", Width: 24},
					{Src: "/images/emoji-star.webp", Alt: "star", Width: 24},
					{Src: "/images/emoji-star.webp", Alt: "star", Width: 24},
					{Src: "/images/emoji-star.webp", Alt: "star", Width: 24},
					{Src: "/images/emoji-star.webp", Alt: "star", Width: 24},
				},
				ValueClass: "",
				Suffix:     "Customer rating 9.0 | 60,259 reviews",
				BoxClass:   "",
				LinkPrefix: "",
				LinkText:   "",
				LinkHref:   "#",
			},
			{
				Label:     "Rated",
				ValueText: "2,307,049",
				ValueImages: []models.RightSideValueImage{
					{Src: "/images/mail.gif", Alt: "mail-gif", Width: 42},
				},
				ValueClass: "right-box-mail",
				Suffix:     "remote job seekers",
				BoxClass:   "",
				LinkPrefix: "",
				LinkText:   "",
				LinkHref:   "#",
			},
			{
				Label:       "Crossposted to",
				ValueText:   "✨ 239 job boards",
				ValueImages: []models.RightSideValueImage{},
				ValueClass:  "",
				Suffix:      "",
				BoxClass:    "",
				LinkPrefix:  "that currently use",
				LinkText:    "our API",
				LinkHref:    "#",
			},
			{
				Label:       "Guaranteed",
				ValueText:   "🎡 200+ clicks",
				ValueImages: []models.RightSideValueImage{},
				ValueClass:  "",
				Suffix:      "or we auto bump it for free",
				BoxClass:    "",
				LinkPrefix:  "",
				LinkText:    "",
				LinkHref:    "#",
			},
			{
				Label:     "Pay with",
				ValueText: "",
				ValueImages: []models.RightSideValueImage{
					{Src: "/images/visa.webp", Alt: "payment-method", Width: 0},
				},
				ValueClass: "",
				Suffix:     "🔐 Secure payment with Stripe",
				BoxClass:   "payment",
				LinkPrefix: "",
				LinkText:   "",
				LinkHref:   "#",
			},
		},
	}
}

func (r *rightSideInfoData) GetAll() []models.RightSideInfo {
	return r.items
}
