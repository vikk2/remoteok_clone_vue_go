package localdata

import "backend/models"

type BenefitData interface {
	GetAll() models.BenefitData
}

type benefitData struct {
	items models.BenefitData
}

func InitBenefitData() BenefitData {
	return &benefitData{
		items: models.BenefitData{
			"💰 401(k)", "🌎 Distributed team", "⏰ Async", "🤓 Vision insurance", "🦷 Dental insurance", "🚑 Medical insurance", "🏖 Unlimited vacation", "🏖 Paid time off", "📆 4 day workweek", "💰 401k matching", "🏔 Company retreats", "🏬 Coworking budget", "📚 Learning budget", "💪 Free gym membership", "🧘 Mental wellness budget", "🖥 Home office budget", "🥧 Pay in crypto", "🥸 Pseudonymous", "💰 Profit sharing", "💰 Equity compensation", "⬜️ No whiteboard interview", "👀 No monitoring system", "🚫 No politics at work", "🎅 We hire old (and young)",
		},
	}
}

func (b *benefitData) GetAll() models.BenefitData {
	return b.items
}
