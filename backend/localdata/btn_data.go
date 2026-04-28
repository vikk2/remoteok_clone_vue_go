package localdata

import (
	"backend/models"
)

type BtnData interface {
	GetAll() []models.BtnData
}

type btnData struct {
	items []models.BtnData
}

func InitBtnData() BtnData {
	return &btnData{
		items: []models.BtnData{
			{
				NavBtns: []models.NavBtn{
					{ID: 1, ClassName: "btn_health", Text: "Health insurance", Link: "#"},
					{ID: 2, ClassName: "btn_post", Text: "Post a job →", Link: "/post-job"},
					{ID: 3, ClassName: "btn_login", Text: "Log in", Link: "/login"},
				},
				PostBtns: []models.PostBtn{
					{ID: 1, ClassName: "btn_post1", Text: "Post a remote job →", Link: "/post-job"},
				},
				HideBtns: []models.HideBtn{
					{ID: 1, ClassName: "btn_hide", Text: "Hide this", IdAttr: "hide"},
				},
				FooterBtns: []models.FooterBtn{
					{
						SubscribeBtn: models.SubscribeBtn{ID: 1, ClassName: "footer-btn", Text: "Subscribe"},
						FooterInput: models.FooterInput{ID: 1, ClassName: "footer-input", Placeholder: "Type your email..."},
					},
				},
			},
		},
	}
}

func (b *btnData) GetAll() []models.BtnData{
	return b.items
}


