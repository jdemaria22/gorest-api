package adapter

import (
	"bytes"
	modeladapter "demoproject/src/adapter/model"
	logging "demoproject/src/log"

	"encoding/json"
	"net/http"
	"os"
)

const (
	CONTENT_TYPE_JSON                = "application/json"
	DEFAULT_PATH_USER_VALUES         = "/stl/userValues"
	DEFAULT_PATH_USER_PENDING_MATCH  = "/stl/getPendingMatchs"
	DEFAULT_PATH_USERS_TASK          = "/stl/userTask"
	DEFAULT_PATH_USERS_TASK_PRIORITY = "/stl/userTaskPriority"
	DEFAULT_PATH_CAC_ORDER_TERMINAL  = "/terminal/cacOrderTerminal"
)

func UserValues(usrId string) (*http.Response, bool) {
	userValuesRequest := modeladapter.UserValuesRequest{UsrId: usrId}
	userValuesRequestJson, err := json.Marshal(userValuesRequest)
	if err != nil {
		logging.Error("::: Error to marshal object userValuesRequestJson :::")
		return nil, false
	}

	url := os.Getenv("URLAPPUSER") + DEFAULT_PATH_USER_VALUES
	logging.Info("::: " + url + " :::")

	resp, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewBuffer(userValuesRequestJson))

	if err != nil {
		logging.Error("::: Error to integrate: " + err.Error() + " :::")
		return nil, false
	}
	return resp, true
}

func UserPendingMatch(usrId string) (*http.Response, bool) {
	userValuesRequest := modeladapter.UserValuesRequest{UsrId: usrId}
	userValuesRequestJson, err := json.Marshal(userValuesRequest)
	if err != nil {
		logging.Error("::: Error to marshal object userValuesRequestJson :::")
		return &http.Response{}, false
	}

	url := os.Getenv("URLAPPUSER") + DEFAULT_PATH_USER_PENDING_MATCH
	logging.Info("::: " + url + " :::")

	resp, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewBuffer(userValuesRequestJson))

	if err != nil {
		logging.Error("::: Error to integrate: " + err.Error() + " :::")
		return &http.Response{}, false
	}
	return resp, true
}

func CacOrderTerminals(cacOrderTerminalsRequest modeladapter.CacOrderTerminalsRequest) (*http.Response, bool) {
	cacOrderTerminalsRequestJson, err := json.Marshal(cacOrderTerminalsRequest)
	if err != nil {
		logging.Error("::: Error to marshal object cacOrderTerminalsRequestJson :::")
		return nil, false
	}

	url := os.Getenv("URLAPPTERMINAL") + DEFAULT_PATH_CAC_ORDER_TERMINAL
	logging.Info("::: " + url + " :::")

	resp, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewBuffer(cacOrderTerminalsRequestJson))
	if err != nil {
		logging.Error("Error to integrate: " + err.Error())
		return nil, false
	}
	return resp, true
}

func UserTask(userTaskRequest modeladapter.UserTaskRequest) (*http.Response, bool) {
	userTaskRequestJson, err := json.Marshal(userTaskRequest)
	if err != nil {
		logging.Error("::: Error to marshal object userTaskRequestJson :::")
		return nil, false
	}

	url := os.Getenv("URLAPPUSER") + DEFAULT_PATH_USERS_TASK
	logging.Info("::: " + url + " :::")

	resp, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewBuffer(userTaskRequestJson))
	if err != nil {
		logging.Error("::: Error to integrate: " + err.Error() + " :::")
		return nil, false
	}
	return resp, true
}

func UserTaskPriority(userTaskRequestPriority modeladapter.UserTaskPriorityRequest) (*http.Response, bool) {
	userTaskRequestPriorityJson, err := json.Marshal(userTaskRequestPriority)
	if err != nil {
		logging.Error("::: Error to marshal object userTaskRequestPriorityJson :::")
		return nil, false
	}

	url := os.Getenv("URLAPPUSER") + DEFAULT_PATH_USERS_TASK_PRIORITY
	logging.Info("::: " + url + " :::")

	resp, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewBuffer(userTaskRequestPriorityJson))
	if err != nil {
		logging.Error("::: Error to integrate: " + err.Error() + " :::")
		return nil, false
	}
	return resp, true
}
