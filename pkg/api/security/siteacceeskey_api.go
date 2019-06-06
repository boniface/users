package security

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/security"
)

const url = api.BASE_URL + "/security/site"

type SiteAccessKeysApi security.SiteAccessKeysApi

func GetSiteAccessKeysApis() ([]SiteAccessKeysApi, error) {
	siteKeys := []SiteAccessKeysApi{}
	resp, _ := api.Rest().Get(url + "/all")
	if resp.IsError() {
		return siteKeys, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &siteKeys)
	if err != nil {
		return siteKeys, errors.New(resp.Status())
	}
	return siteKeys, nil

}

func getSiteAccessKeysApi(id string) (SiteAccessKeysApi, error) {
	siteKey := SiteAccessKeysApi{}
	resp, _ := api.Rest().Get(url + "/get/" + id)
	if resp.IsError() {
		return siteKey, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &siteKey)
	if err != nil {
		return siteKey, errors.New(resp.Status())
	}
	return siteKey, nil

}

func createSiteAccessKeysApi(entity SiteAccessKeysApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateSiteAccessKeysApi(entity SiteAccessKeysApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteSiteAccessKeysApi(entity SiteAccessKeysApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
