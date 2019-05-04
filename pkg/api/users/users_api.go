package users

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/users"
)

const userurl = api.BASE_URL + "/users"

type User users.User

func getUsers() ([]User, error) {
	users := []User{}
	resp, _ := api.Rest().Get(userurl + "/all")
	if resp.IsError() {
		return users, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &users)
	if err != nil {
		return users, errors.New(resp.Status())
	}
	return users, nil

}

func getUser(id string) (User, error) {
	user := User{}
	resp, _ := api.Rest().Get(userurl + "/get/" + id)
	if resp.IsError() {
		return user, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &user)
	if err != nil {
		return user, errors.New(resp.Status())
	}
	return user, nil

}

func createUser(entity User) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateUser(entity User) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteUser(entity User) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(userurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
