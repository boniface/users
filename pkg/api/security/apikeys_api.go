package security

import (
	"errors"
	"fmt"
	"users/pkg/api"
	"users/pkg/domain/security"
)

const apiurl = api.BASE_URL + "/security"

type ApiKeys security.ApiKeys

func GetApiKeys() ([]ApiKeys, error) {
	apikeys := []ApiKeys{}
	resp, _ := api.Rest().Get(apiurl + "/all")
	if resp.IsError() {
		return apikeys, errors.New(resp.Status())
	}

	err := api.JSON.Unmarshal(resp.Body(), &apikeys)
	if err != nil {
		fmt.Println(" There is an Error ", err)
		return apikeys, errors.New(resp.Status())
	}
	return apikeys, nil

}

func getApiKey(id string) (ApiKeys, error) {
	apikey := ApiKeys{}
	resp, _ := api.Rest().Get(apiurl + "/get/" + id)
	if resp.IsError() {
		return apikey, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &apikey)
	if err != nil {
		return apikey, errors.New(resp.Status())
	}
	return apikey, nil

}

func createApiKey(entity ApiKeys) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(apiurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
func updateApiKey(entity ApiKeys) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(apiurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func deleteApiKey(entity ApiKeys) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(apiurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
