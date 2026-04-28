package models

type FeedbackData struct {
	ID int `json:"id"`
	Profile_img string `json:"profile_img"`
	Icon string `json:"icon,omitempty"`
	Feedback string `json:"feedback"`
}