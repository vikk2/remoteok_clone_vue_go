package localdata

import "backend/models"

type FeedbackData interface {
	GetAll() []models.FeedbackData
}

type feedbackData struct {
	items []models.FeedbackData
}

func InitFeedbackData() FeedbackData {
	return &feedbackData{
		items: []models.FeedbackData{
			{ID: 1, Profile_img: "/images/pf-1.webp", Icon: "/images/green-circle.webp", Feedback: "<span>\"The response has been amazing! A lotttt of applicants.</span> Thank you, and everything Remote OK did, to help with this!\"<br><br> — <a href=\"#\">Edwin</a>, DeRel at <a href=\"#\">Open AI</a>"},
			{ID: 2, Profile_img: "/images/pf-2.jpg", Icon: "/images/green-circle.webp", Feedback: "\"Remote OK has been <span>an essential platform</span> for attracting great talent to our remote-first company\"<br><br> — <a href=\"#\">Sara</a>, Recruitment at <a href=\"#\">Komoot</a>"},
			{ID: 3, Profile_img: "/images/pf-3.webp", Icon: "/images/icon-red.webp", Feedback: "\"I want to say that <span>having 1,500 applicants to sort through</span> is a good problem to have and thanks to y'all for making that happen\"<br><br> — <a href=\"#\">Sara</a>, Recruitment at <a href=\"#\">Komoot</a>"},
			{ID: 4, Profile_img: "/images/pf-4.webp", Icon: "/images/icon-red.webp", Feedback: "\"FYI - <span>We are loving the performance of the job ads,</span> we have had so many great applicants. Thank you for your great work on building this job board :)\"<br><br> — <a href\"#\">Tris</a>, Global Sourcing Lead at <a href=\"#\">Aula Education</a>"},
			{ID: 5, Profile_img: "/images/pf-5.webp", Icon: "/images/pf-5.webp", Feedback: "\"I love the site, I tried so many different job sites and <span>Remote OK is by far the best to deliver</span>\"<br><br> — <a href=\"#\">Zsolt</a>, Founder & CEO at <a href=\"#\">PingPong</a>"},
			{ID: 6, Profile_img: "/images/pf-6.webp", Icon: "/images/icon-red.webp", Feedback: "\"We super like Remote OK. <span>We had around 100 applicants</span> for our previous job post, so we are doing a new one :)\"<br><br> — <a href=\"#\">Baptiste</a>, Founder & CEO at <a href=\"#\">Crisp Chat</a>"},
			{ID: 7, Profile_img: "/images/pf-7.webp", Icon: "/images/icon-red.webp", Feedback: "\"Awesome, you rock! <span>Your customer support is a perfect example of why I stick with your site</span> as our \"go-to\" job board...\"<br><br> — <a href=\"#\">Ken</a>, CEO at <a href=\"#\">Savvy</a>"},
			{ID: 8, Profile_img: "/images/pf-8.webp", Icon: "/images/icon-red.webp", Feedback: "Hey Pieter, just a quick shout out to your amazing stuff at Remote OK - we posted 2 jobs in the recent month now, and a quick feedback on the applicants: we not just getting a ton of them, but <span>people coming from Remote OK are much much much more valid and higher quality</span> compared to people coming from LinkedIn, Angel, VirtualVocations, Remotive, or Indeed, as we posted our jobs there too. So big fan of Remote OK, the last job was only posted there. Thanks for the good stuff!<br><br> — <a href=\"#\">Peter</a>, COO at <a href=\"#\">KISSPatent</a>"},
			{ID: 9, Profile_img: "/images/pf-9.webp", Icon: "/images/icon-red.webp", Feedback: "\"Thank you for the collaboration, it was a pleasure working with you and with Remote OK. It was <span>the best job board targeted at remote professionals in terms of results.</span>\"<br><br> — <a href=\"#\" class=\"a-gray\">Alex</a>, Global Sourcing Strategist at <a href=\"#\">Crossover</a>"},
			{ID: 10, Profile_img: "/images/pf-10.webp", Feedback: "👋 Hi! I'm the maker of Remote OK and other sites related to remote work such as <a href=\"#\">Nomads.com</a>. Remote OK isn't a big team, it's actually just a one-man operation which is me on a laptop somewhere in the world. I built Remote OK to help 🚀 accelerate the revolution that is remote work (and pay my rent and coffee). My site has now become the <a href=\"#\">#1 remote job platform</a> in the world!<br><br>Remote work gives people more flexibility in their daily lives, and lets employers hire the 🏆 best talent regardless of where they are located in the 🌍 world. That's the most significant change to work since the industrial revolution, and that's why I make this site.<br><br>Need help posting your job? Here's <a href=\"#\">my 𝕏</a>. Tweet to me and I'll help you personally and can also help you if you'd like to buy multiple <a href=\"#\">job posts packages</a> or <a href=\"#\">get a discount</a>.<br><br>Go remote!<br><br> — <a href=\"#\">Pieter Levels</a>, Solofounder of <a href=\"#\" class=\"a-gray\">Remote OK</a>"},
		},
	}
}

func (f *feedbackData) GetAll() []models.FeedbackData {
	return f.items
}
