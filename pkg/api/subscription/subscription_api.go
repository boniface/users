package subscription

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/mail"
)

const url = api.BASE_URL + "/roles"

type Entity mail.MailApi

func  getEntities() ([]Entity, error){
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

func getEntity(entitId string ) (Entity, error){
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

func createEntity(entity Entity ) (bool, error){
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateEntity(entity Entity ) (bool, error){
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteEntity(entity Entity ) (bool, error){
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
