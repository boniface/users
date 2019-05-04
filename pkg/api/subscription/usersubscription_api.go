package subscription

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/subscription"
)

const url = api.BASE_URL + "/subscriptions/user"

type UserSubscriptions subscription.UserSubscriptions

func getUserSubscriptions() ([]UserSubscriptions, error) {
	userSubscriptions := []UserSubscriptions{}
	resp, _ := api.Rest().Get(url + "/all")
	if resp.IsError() {
		return userSubscriptions, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userSubscriptions)
	if err != nil {
		return userSubscriptions, errors.New(resp.Status())
	}
	return userSubscriptions, nil

}

func getUserSubscription(id string) (UserSubscriptions, error) {
	userSubscriptions := UserSubscriptions{}
	resp, _ := api.Rest().Get(url + "/get/" + id)
	if resp.IsError() {
		return userSubscriptions, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &userSubscriptions)
	if err != nil {
		return userSubscriptions, errors.New(resp.Status())
	}
	return userSubscriptions, nil

}

func createUserSubscription(entity UserSubscriptions) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateUserSubscription(entity UserSubscriptions) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func deleteUserSubscription(entity UserSubscriptions) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
