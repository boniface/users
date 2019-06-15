package subscription

import (
	"errors"
	"users/pkg/api"
	"users/pkg/domain/subscription"
)

const url = api.BASE_URL + "/subscriptions/user"

type UserSubscriptions subscription.UserSubscriptions

func GetUserSubscriptions() ([]UserSubscriptions, error) {
	userSubscriptions := []UserSubscriptions{}
	resp, _ := api.Rest().Get(url + "/get/all")
	if resp.IsError() {
		return userSubscriptions, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &userSubscriptions)
	if err != nil {
		return userSubscriptions, errors.New(resp.Status())
	}
	return userSubscriptions, nil

}

func GetUserSubscription(id string) (UserSubscriptions, error) {
	userSubscriptions := UserSubscriptions{}
	resp, _ := api.Rest().Get(url + "/get/" + id)
	if resp.IsError() {
		return userSubscriptions, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &userSubscriptions)
	if err != nil {
		return userSubscriptions, errors.New(resp.Status())
	}
	return userSubscriptions, nil

}

func CreateUserSubscription(entity UserSubscriptions) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func UpdateUserSubscription(entity UserSubscriptions) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func DeleteUserSubscription(entity UserSubscriptions) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
