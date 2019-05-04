package login

import "time"

type LoginCredential struct {
	Email    string `json:"email"`
	SiteId   string `json:"siteId"`
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type LoginStatus struct {
	Status string `json:"status"`
}

type PasswordChangeCredentials struct {
	Email       string `json:"email"`
	UserId      string `json:"userId"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type Profile struct {
	Email  string `json:"email"`
	SiteId string `json:"siteId"`
	UserId string `json:"userId"`
}

type ResetToken struct {
	Resetokenvalue string `json:"resetokenvalue"`
	UserId         string `json:"userId"`
}

type UserLoginToken struct {
	SiteId     string    `json:"siteId"`
	UserId     string    `json:"userId"`
	Id         string    `json:"id"`
	ExpiryDate time.Time `json:"expiryDate"`
	TokenValue string    `json:"tokenValue"`
	Message    string    `json:"message"`
}
