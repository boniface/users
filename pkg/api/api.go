package api

import (
	"gopkg.in/resty.v1"
)

const BASE_URL string = "https://hashusersapi.hash-code.com"

func Rest() *resty.Request{
	return resty.R().SetAuthToken("").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "application/json")
}


