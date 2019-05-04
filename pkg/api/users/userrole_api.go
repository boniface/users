package users

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/users"
)

const userroleurl = api.BASE_URL + "/users/role"

type UserRole users.UserRole

func getUserRoles() ([]UserRole, error) {
	userRoles := []UserRole{}
	resp, _ := api.Rest().Get(userroleurl + "/all")
	if resp.IsError() {
		return userRoles, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userRoles)
	if err != nil {
		return userRoles, errors.New(resp.Status())
	}
	return userRoles, nil

}

func getUserRole(id string) (UserRole, error) {
	userRole := UserRole{}
	resp, _ := api.Rest().Get(userroleurl + "/get/" + id)
	if resp.IsError() {
		return userRole, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userRole)
	if err != nil {
		return userRole, errors.New(resp.Status())
	}
	return userRole, nil

}

func createUserRole(entity UserRole) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userroleurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateUserRole(entity UserRole) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userroleurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteUserRole(entity UserRole) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userroleurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
