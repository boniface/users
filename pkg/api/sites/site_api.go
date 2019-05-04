package sites

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/sites"
)

const url = api.BASE_URL + "/sites"

type Site sites.Site

func getSites() ([]Site, error) {
	entities := []Site{}
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

func getSite(id string) (Site, error) {
	role := Site{}
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

func createSite(entity Site) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateSite(entity Site) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteSite(entity Site) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
