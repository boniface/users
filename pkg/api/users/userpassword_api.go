package users

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/users"
)

const usersurl = api.BASE_URL + "/users/pass"

type UserPassword users.UserPassword

func getUserPasswords() ([]UserPassword, error) {
	userPasswords := []UserPassword{}
	resp, _ := api.Rest().Get(usersurl + "/all")
	if resp.IsError() {
		return userPasswords, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userPasswords)
	if err != nil {
		return userPasswords, errors.New(resp.Status())
	}
	return userPasswords, nil
}

func getUserPassword(id string) (UserPassword, error) {
	userPassword := UserPassword{}
	resp, _ := api.Rest().Get(usersurl + "/get/" + id)
	if resp.IsError() {
		return userPassword, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userPassword)
	if err != nil {
		return userPassword, errors.New(resp.Status())
	}
	return userPassword, nil
}

func createUserPassword(entity UserPassword) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func updateUserPassword(entity UserPassword) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func deleteUserPassword(entity UserPassword) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
