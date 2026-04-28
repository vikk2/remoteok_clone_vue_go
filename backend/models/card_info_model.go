package models

type CardDetails struct {
	JobDescription string `json:"job_description"`
	WebsiteURL     string `json:"websiteURL"`
	QRCode         string `json:"qr_code"`
}

type CardInfo struct {
	Type          string       `json:"type"`
	JobClass      string       `json:"jobClass"`
	CardVariant   string       `json:"cardVariant,omitempty"`
	LogoType      string       `json:"logoType,omitempty"`
	Logo          string       `json:"logo,omitempty"`
	LogoText      string       `json:"logoText,omitempty"`
	JobTitle      string       `json:"jobTitle,omitempty"`
	Company       string       `json:"company,omitempty"`
	Location      string       `json:"location,omitempty"`
	Salary        string       `json:"salary,omitempty"`
	Schedule      string       `json:"schedule,omitempty"`
	New           bool         `json:"new,omitempty"`
	Verify        bool         `json:"verify,omitempty"`
	BalloonIcon   bool         `json:"balloonIcon,omitempty"`
	Label         []string     `json:"label,omitempty"`
	Date          string       `json:"date,omitempty"`
	Details       *CardDetails `json:"details,omitempty"`
	AdTitle       string       `json:"adTitle,omitempty"`
	AdDescription string       `json:"adDescription,omitempty"`
	AdImage       string       `json:"adImage,omitempty"`
}
