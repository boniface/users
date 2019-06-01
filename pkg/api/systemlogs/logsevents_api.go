package systemlogs

import (
	"encoding/json"
	"errors"
	"users/pkg/api"
	"users/pkg/domain/systemlogs"
)

const url = api.BASE_URL + "/systemlogs"

type LogEvent systemlogs.LogEvent

func GetLogEvents() ([]LogEvent, error) {
	logs := []LogEvent{}
	resp, _ := api.Rest().Get(url + "/all")
	if resp.IsError() {
		return logs, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &logs)
	if err != nil {
		return logs, errors.New(resp.Status())
	}
	return logs, nil

}

func GetLogEventsForSite(siteid string) ([]LogEvent, error) {
	logs := []LogEvent{}
	resp, _ := api.Rest().Get(url + "/get/site/" + siteid)
	if resp.IsError() {
		return logs, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &logs)
	if err != nil {
		return logs, errors.New(resp.Status())
	}
	return logs, nil

}

func GetLogEvent(id string) (LogEvent, error) {

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

func CreateLogEvent(entity LogEvent) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}
func UpdateLogEvent(entity LogEvent) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func DeleteLogEvent(entity LogEvent) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(url + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
