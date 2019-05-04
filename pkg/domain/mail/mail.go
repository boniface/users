package mail

import "time"

type EmailMessage struct {
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

type MailApi struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Sender string `json:"sender"`
}

type MailConfig struct {
	SiteId string    `json:"siteId"`
	Id     string    `json:"id"`
	Key    string    `json:"key"`
	Value  string    `json:"value"`
	Host   string    `json:"host"`
	Port   string    `json:"port"`
	State  string    `json:"state"`
	Date   time.Time `json:"date"`
}

type MessageResponse struct {
	StatusCode string            `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

type SmtpConfig struct {
	Id       string `json:"id"`
	Port     string `json:"port"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}