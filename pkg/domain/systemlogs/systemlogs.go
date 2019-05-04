package systemlogs

type LogEvent struct {
	SiteId    string `json:"siteId"`
	Id        string `json:"id"`
	EventName string `json:"eventName"`
	EventType string `json:"eventType"`
	Message   string `json:"message"`
	Date      string `json:"date"`
}
