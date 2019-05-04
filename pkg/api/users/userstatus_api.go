package users

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/users"
)

const userstatusurl = api.BASE_URL + "/users/status"

type UserStatus users.UserStatus

func getUserStatuss() ([]UserStatus, error) {
	userStatuses := []UserStatus{}
	resp, _ := api.Rest().Get(userstatusurl + "/all")
	if resp.IsError() {
		return userStatuses, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userStatuses)
	if err != nil {
		return userStatuses, errors.New(resp.Status())
	}
	return userStatuses, nil

}

func getUserStatus(id string) (UserStatus, error) {
	userStatus := UserStatus{}
	resp, _ := api.Rest().Get(userstatusurl + "/get/" + id)
	if resp.IsError() {
		return userStatus, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userStatus)
	if err != nil {
		return userStatus, errors.New(resp.Status())
	}
	return userStatus, nil

}

func createUserStatus(entity UserStatus) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userstatusurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateUserStatus(entity UserStatus) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(userstatusurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteUserStatus(entity UserStatus) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(userstatusurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
