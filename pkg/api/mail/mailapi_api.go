package mail

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/mail"
)

var mailurl = api.BASE_URL + "/mail/api"

type MailApi mail.MailApi

func getMailApis() ([]MailApi, error) {
	entities := []MailApi{}
	resp, _ := api.Rest().Get(mailurl + "/all")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil

}

func getMailApi(id string) (MailApi, error) {
	entity := MailApi{}
	resp, _ := api.Rest().Get(mailurl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func createMailApi(MailApi MailApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(MailApi).
		Post(mailurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateMailApi(MailApi MailApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(MailApi).
		Post(mailurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteMailApi(MailApi MailApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(MailApi).
		Post(mailurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
