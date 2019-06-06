package mail

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/mail"
)

var url = api.BASE_URL + "/mail/smtp"

type SmtpConfig mail.SmtpConfig

func GetSmtpConfigs() ([]SmtpConfig, error) {
	entities := []SmtpConfig{}
	resp, _ := api.Rest().Get(url + "/all")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil

}

func getSmtpConfig(id string) (SmtpConfig, error) {
	role := SmtpConfig{}
	resp, _ := api.Rest().Get(url + "/get/" + id)
	if resp.IsError() {
		return role, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &role)
	if err != nil {
		return role, errors.New(resp.Status())
	}
	return role, nil

}

func createEntity(entity SmtpConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
func updateEntity(entity SmtpConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func deleteEntity(entity SmtpConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
