package subscription

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/subscription"
)

const siteSubs = api.BASE_URL + "/subscriptions/site"

type SiteSubscription subscription.SiteSubscription

func GetSiteSubscriptions() ([]SiteSubscription, error) {
	subscriptions := []SiteSubscription{}
	resp, _ := api.Rest().Get(siteSubs + "/all")
	if resp.IsError() {
		return subscriptions, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &subscriptions)
	if err != nil {
		return subscriptions, errors.New(resp.Status())
	}
	return subscriptions, nil

}

func GetSiteSubscription(id string) (SiteSubscription, error) {
	subscription := SiteSubscription{}
	resp, _ := api.Rest().Get(siteSubs + "/get/" + id)
	if resp.IsError() {
		return subscription, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &subscription)
	if err != nil {
		return subscription, errors.New(resp.Status())
	}
	return subscription, nil

}

func CreateSiteSubscription(entity SiteSubscription) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(siteSubs + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func UpdateSiteSubscription(entity SiteSubscription) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(siteSubs + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func DeleteSiteSubscription(entity SiteSubscription) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(siteSubs + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
