package mail

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/mail"
)

var  mailconfig = api.BASE_URL + "/roles"

type MailConfig mail.MailConfig

func  getEntities() ([]MailConfig, error){
	roles:=[]Role{}
	resp, _ := api.Rest().Get(mailconfig + "/all")
	if resp.IsError() {
		return roles , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &roles)
	if err != nil {
		return roles ,  errors.New(resp.Status())
	}
	return roles , nil

}

func getEntity(entitId string ) (MailConfig, error){
	role := Role{}
	resp, _ := api.Rest().Get(mailconfig + "/get/" + id)
	if resp.IsError() {
		return role , errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &role)
	if err != nil {
		return role ,  errors.New(resp.Status())
	}
	return role , nil

}

func createEntity(entity MailConfig) (bool, error){

}
func updateEntity(entity MailConfig) (bool, error){

}

func deleteEntity(entity MailConfig) (bool, error){

}