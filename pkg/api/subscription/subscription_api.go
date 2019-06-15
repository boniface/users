package subscription

import (
	"errors"
	"users/pkg/api"
	"users/pkg/domain/subscription"
)

const subsurl = api.BASE_URL + "/subscriptions"

type Subscription subscription.Subscription

func GetSubscriptions() ([]Subscription, error) {
	subscriptions := []Subscription{}
	resp, _ := api.Rest().Get(subsurl + "/getall")
	if resp.IsError() {
		return subscriptions, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &subscriptions)
	if err != nil {
		return subscriptions, errors.New(resp.Status())
	}
	return subscriptions, nil

}

func GetSubscription(id string) (Subscription, error) {
	subscription := Subscription{}
	resp, _ := api.Rest().Get(subsurl + "/get/" + id)
	if resp.IsError() {
		return subscription, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &subscription)
	if err != nil {
		return subscription, errors.New(resp.Status())
	}
	return subscription, nil
}

func CreateSubscription(entity Subscription) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subsurl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
func UpdateSubscription(entity Subscription) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subsurl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func DeleteSubscription(entity Subscription) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subsurl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
