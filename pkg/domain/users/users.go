package users

import "time"

type Account struct {
	SiteId   string `json:"siteId"`
	Email    string `json:"email"`
	UserId   string `json:"userId"`
	SiteName string `json:"siteName"`
}

type Registration struct {
	Email  string `json:"email"`
	SiteId string `json:"siteId"`
}

type User struct {
	SiteId      string    `json:"siteId"`
	Email       string    `json:"email"`
	UserId      string    `json:"userId"`
	ScreenName  string    `json:"screenName"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
	Password    string    `json:"password"`
	DateCreated time.Time `json:"dateCreated"`
}

type UserPassword struct {
	UserId   string `json:"userId"`
	Date     string `json:"date"`
	Password string `json:"password"`
}

type UserRole struct {
	UserId string `json:"userId"`
	Date   string `json:"date"`
	RoleId string `json:"roleId"`
}

type UserStatus struct {
	UserId string `json:"userId"`
	Date   string `json:"date"`
	Status string `json:"status"`
}
