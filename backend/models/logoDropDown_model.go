package models

type SectionItem struct {
	Title string      `json:"title,omitempty"`
	Items []ItemsData `json:"items"`
}

type ItemsData struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Link  string `json:"href"`
	Image string `json:"img,omitempty"`
}


