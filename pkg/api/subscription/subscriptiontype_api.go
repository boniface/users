package subscription

import (
	"errors"
	"users/pkg/api"
	"users/pkg/domain/subscription"
)

const subtypesurl = api.BASE_URL + "/subscriptions/types"

type SubscriptionTypes subscription.SubscriptionTypes

func GetSubscriptionTypes() ([]SubscriptionTypes, error) {
	roles := []SubscriptionTypes{}
	resp, _ := api.Rest().Get(subtypesurl + "/get")
	if resp.IsError() {
		return roles, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &roles)
	if err != nil {
		return roles, errors.New(resp.Status())
	}
	return roles, nil

}

func GetSubscriptionType(id string) (SubscriptionTypes, error) {
	role := SubscriptionTypes{}
	resp, _ := api.Rest().Get(subtypesurl + "/get/" + id)
	if resp.IsError() {
		return role, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &role)
	if err != nil {
		return role, errors.New(resp.Status())
	}
	return role, nil

}

func CreateSubscriptionType(entity SubscriptionTypes) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subtypesurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func UpdateSubscriptionType(entity SubscriptionTypes) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subtypesurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func DeleteSubscriptionType(entity SubscriptionTypes) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subtypesurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
