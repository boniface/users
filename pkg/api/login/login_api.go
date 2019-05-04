package login

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/login"
	"users/pkg/domain/mail"
	"users/pkg/domain/security"
	"users/pkg/domain/users"
)


const roleurl = api.BASE_URL + "/login"
type User users.User
type MessageResponse mail.MessageResponse
type UserGeneratedToken security.UserGeneratedToken
type Account users.Account
type Profile login.Profile
type LoginCredential login.LoginCredential

func resetPasswordRequest(resetKey string) (MessageResponse, error) {
	entity := MessageResponse{}
	resp, _ := api.Rest().Get(roleurl + "/get/" + resetKey)
	if resp.IsError() {
		return entity , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity ,  errors.New(resp.Status())
	}
	return entity , nil

}

func isUserRegistered(user User) (bool, error) {
	resp, _ := api.Rest().
		SetBody(user).
		Post(roleurl + "/registered")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func forgotPassword(profile Profile) (MessageResponse, error) {
	entity := MessageResponse{}
	resp, _ := api.Rest().
		SetBody(profile).
		Post(roleurl + "/forgotpassword")
	if resp.IsError() {
		return entity , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity ,  errors.New(resp.Status())
	}
	return entity , nil

}

func getUserAccounts(email string) ([] Account, error) {
	entities:=[]Account{}
	resp, _ := api.Rest().Get(roleurl + "/get/"+email)
	if resp.IsError() {
		return entities , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities ,  errors.New(resp.Status())
	}
	return entities , nil

}

func getLoginToken(loginCredential LoginCredential) (UserGeneratedToken, error) {
	entity := UserGeneratedToken{}
	resp, _ := api.Rest().
		SetBody(loginCredential).
		Post(roleurl + "/getlogintoken")
	if resp.IsError() {
		return entity , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity ,  errors.New(resp.Status())
	}
	return entity , nil

}

func isUserLoggedIn(user User) (bool, error) {
	// GET THIS FROM SESSION
	//resp, _ := api.Rest().
	//	SetBody(entity).
	//	Post(roleurl + "/create")
	//if resp.IsError() {
	//	return false, errors.New(resp.Status())
	//}

	return true, nil

}

func getUserRole(user User) (string, error) {
	// GET THIS FROM THE SESSION

	//role := Role{}
	//resp, _ := api.Rest().Get(roleurl + "/get/" + id)
	//if resp.IsError() {
	//	return role , errors.New(resp.Status())
	//}
	//err := json.Unmarshal(resp.Body(), &role)
	//if err != nil {
	//	return role ,  errors.New(resp.Status())
	//}
	return "" , nil

}

func logout() (bool, error) {
	// INVALIDATE THE SESSION

	//resp, _ := api.Rest().
	//	SetBody(entity).
	//	Post(roleurl + "/create")
	//if resp.IsError() {
	//	return false, errors.New(resp.Status())
	//}

	return true, nil

}
