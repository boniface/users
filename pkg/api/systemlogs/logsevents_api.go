package systemlogs

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/systemlogs"
)

const url = api.BASE_URL + "/roles"

type LogEvent systemlogs.LogEvent

func getLogEvents() ([]LogEvent, error) {
	roles := []LogEvent{}
	resp, _ := api.Rest().Get(url + "/all")
	if resp.IsError() {
		return roles, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &roles)
	if err != nil {
		return roles, errors.New(resp.Status())
	}
	return roles, nil

}

func getLogEvent(id string) (LogEvent, error) {

	logEvent := LogEvent{}
	resp, _ := api.Rest().Get(url + "/get/" + id)
	if resp.IsError() {
		return logEvent, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &logEvent)
	if err != nil {
		return logEvent, errors.New(resp.Status())
	}
	return logEvent, nil
}

func createLogEvent(entity LogEvent) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func updateLogEvent(entity LogEvent) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func deleteLogEvent(entity LogEvent) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
