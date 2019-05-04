package mail

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/mail"
)

var mailconfig = api.BASE_URL + "/mail"

type MailConfig mail.MailConfig

func getMailconfigs() ([]MailConfig, error) {
	entities := []MailConfig{}
	resp, _ := api.Rest().Get(mailconfig + "/all")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil

}

func getMailConfig(id string) (MailConfig, error) {
	entity := MailConfig{}
	resp, _ := api.Rest().Get(mailconfig + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func createMailConfig(entity MailConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
func updateMailConfig(entity MailConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func deleteMailConfig(entity MailConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
