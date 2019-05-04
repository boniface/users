package subscription

import "time"

type SiteSubscription struct {
	SiteId         string    `json:"siteId"`
	Id             string    `json:"id"`
	SubscriptionId string    `json:"subscriptionId"`
	TokenValue     string    `json:"tokenValue"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
}

type Subscription struct {
	Id               string    `json:"id"`
	SubscriptionType string    `json:"subscriptionType"`
	Value            float64   `json:"value"`
	Duration         string    `json:"duration"`
	Description      string    `json:"description"`
	DateCreated      time.Time `json:"dateCreated"`
	Status           string    `json:"status"`
}

type SubscriptionTypes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserSubscriptions struct {
	SiteId         string    `json:"siteId"`
	UserId         string    `json:"userId"`
	Id             string    `json:"id"`
	SubscriptionId string    `json:"subscriptionId"`
	TokenValue     string    `json:"tokenValue"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
}