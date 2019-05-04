package roles

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/roles"
)

const url = api.BASE_URL + "/roles"

type Role roles.Role

func GetRoles() ([]Role, error) {
	roles:=[]Role{}
	resp, _ := api.Rest().Get(url + "/all")
	if resp.IsError() {
		return roles , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &roles)
	if err != nil {
		return roles ,  errors.New(resp.Status())
	}
	return roles , nil
}

func GetRole(id string) (Role, error) {
	role := Role{}
	resp, _ := api.Rest().Get(url + "/get/" + id)
	if resp.IsError() {
		return role , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &role)
	if err != nil {
		return role ,  errors.New(resp.Status())
	}
	return role , nil
}

func GetRolesForSite(site string) ([]Role, error) {
	roles:=[]Role{}
	resp, _ := api.Rest().Get(url + "/all")
	if resp.IsError() {
		return roles , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &roles)
	if err != nil {
		return roles ,  errors.New(resp.Status())
	}
	return roles , nil
}

func CreateRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteRole(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
