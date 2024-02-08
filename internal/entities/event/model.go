package event

type EventModel struct {
	OrderType  string `json:"orderType"  example:"CardVerify"`
	SessionId  string `json:"sessionId"  example:"500cf308-e666-4639-aa9f-f6376015d1b4"`
	Card       string `json:"card"       example:"4433**1409"`
	EventDate  string `json:"eventDate"  example:"2023-04-07 05:29:54.362216 +00:00"`
	WebsiteUrl string `json:"websiteUrl" example:"https://adidas.com"`
}
