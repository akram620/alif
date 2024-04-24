package models

type Event struct {
	ID         int64  `json:"id"`
	OrderType  string `json:"orderType"`
	SessionID  string `json:"sessionId"`
	Card       string `json:"card"`
	EventDate  string `json:"eventDate"`
	WebsiteURL string `json:"websiteUrl"`
}
