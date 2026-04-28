package models

type RightSideValueImage struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
	Width int `json:"width,omitempty"`
}

type RightSideInfo struct {
	Label string `json:"label"`
	ValueText string `json:"value_text"`
	ValueImages []RightSideValueImage `json:"value_images"`
	ValueClass string `json:"value_class"`
	Suffix string `json:"suffix"`
	BoxClass string `json:"box_class"`
	LinkPrefix string `json:"link_prefix"`
	LinkText string `json:"link_text"`
	LinkHref string `json:"link_href"`
}
