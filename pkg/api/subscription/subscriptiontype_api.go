package subscription

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/subscription"
)

const subtypesurl = api.BASE_URL + "/subscriptions/types"

type SubscriptionTypes subscription.SubscriptionTypes

func getSubscriptionTypes() ([]SubscriptionTypes, error) {
	roles := []SubscriptionTypes{}
	resp, _ := api.Rest().Get(subtypesurl + "/all")
	if resp.IsError() {
		return roles, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &roles)
	if err != nil {
		return roles, errors.New(resp.Status())
	}
	return roles, nil

}

func getSubscriptionType(id string) (SubscriptionTypes, error) {
	role := SubscriptionTypes{}
	resp, _ := api.Rest().Get(subtypesurl + "/get/" + id)
	if resp.IsError() {
		return role, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &role)
	if err != nil {
		return role, errors.New(resp.Status())
	}
	return role, nil

}

func createSubscriptionType(entity SubscriptionTypes) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subtypesurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateSubscriptionType(entity SubscriptionTypes) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subtypesurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteSubscriptionType(entity SubscriptionTypes) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subtypesurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
