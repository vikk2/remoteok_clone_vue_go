package models

type NavBtn struct {
	ID    int    `json:"id"`
	ClassName string `json:"className"`
	Text string `json:"text"`
	Link string `json:"href"`
}

type PostBtn struct {
	ID    int    `json:"id"`
	ClassName string `json:"className"`
	Text string `json:"text"`
	Link string `json:"href"`
}

type HideBtn struct {
	ID int `json:"id"`
	ClassName string `json:"className"`
	Text string `json:"text"`
	IdAttr string `json:"id_attribute"`
}

type SubscribeBtn struct {
	ID int `json:"id"`
	ClassName string `json:"className"`
	Text string `json:"text"`
}

type FooterInput struct {
	ID int `json:"id"`
	ClassName string `json:"className"`
	Placeholder string `json:"placeHolder"`	
}

type FooterBtn struct {
	SubscribeBtn SubscribeBtn `json:"subscribeBtn"`
	FooterInput FooterInput `json:"footerInput"` 
}

type BtnData struct {
	NavBtns []NavBtn `json:"nav_btns"`
	PostBtns []PostBtn `json:"post_btn"`
	HideBtns []HideBtn `json:"hide_btn"`
	FooterBtns []FooterBtn `json:"footer_btn"`
}


