package security

import "time"

type ApiKeys struct {
	Id     string    `json:"id"`
	Value  string    `json:"value"`
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

type SiteAccessKeysApi struct {
	SiteId     string `json:"siteId"`
	Id         string `json:"id"`
	ExpiryDate string `json:"expiryDate"`
	Value      string `json:"value"`
	Message    string `json:"message"`
}

type UserGeneratedToken struct {
	Token   string `json:"token"`
	Status  string `json:"status"`
	Message string `json:"message"`
	SiteId  string `json:"siteId"`
}