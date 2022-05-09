package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	adapter "demoproject/src/adapter"
	modeladapter "demoproject/src/adapter/model"
	logging "demoproject/src/log"
	"demoproject/src/model"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_DESCRIPTION   = "Error description."
	SUCCESS_DESCRIPTION = "Success description."
	ERROR_STATUS        = 1
	SUCCESS_STATUS      = 0
	INTERMEDIATE        = "I"
	ENTEGADA            = "E"
	ESTADOH             = "H"
)

var (
	wg sync.WaitGroup
)

// user godoc
// @Summary Show user
// @Description Obtains info of Users
// @Tags user
// @Accept  json
// @Produce  json
// @Param  usrId path string true "ID Del usuario"
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Success 200 {object} model.Status
// @Router /user/{usrId} [get]
func GetUser(context *gin.Context) {
	context.Writer.Header().Add("Content-Type", "application/json")
	var user model.User
	usrId := context.Param("usrId")

	resp, ok := adapter.UserValues(usrId)
	if !ok {
		logging.Error("::: Error in adapter.UserValues :::")
		context.JSON(http.StatusInternalServerError, nil)
		return
	}
	if resp.StatusCode == 204 || resp.StatusCode == 404 {
		logging.Info("::: Not found userValules :::")
		context.JSON(http.StatusNoContent, nil)
		return
	}
	if resp.StatusCode >= 500 {
		logging.Error("::: Error in api userValules :::")
		context.JSON(http.StatusNoContent, nil)
		return
	}
	userValues, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error("::: Error to read userValues response :::")
		context.JSON(http.StatusInternalServerError, nil)
		return
	}
	json.Unmarshal(userValues, &user)

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, ok := adapter.UserPendingMatch(usrId)
		if !ok {
			logging.Info("::: Error in adapter UserPendingMatch :::")
			return
		}
		if resp.StatusCode != 200 {
			return
		}
		userPendingMatch, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logging.Error("::: Error to read UserPendingMatch adapter :::")
			return
		}

		type responsePendingMatch struct {
			TerminalPendingMatchs   int    `json:"terminalPendingMatchs"`
			AccesoriesPendingMatchs int    `json:"accesoriesPendingMatchs"`
			ErrorCode               int    `json:"errorCode"`
			ErrorDescription        string `json:"errorDescription"`
		}
		respPendingMatch := responsePendingMatch{}
		json.Unmarshal(userPendingMatch, &respPendingMatch)
		if respPendingMatch.ErrorCode != 0 {
			return
		}
		user.AccountingItemsInfo.PendingMaterials = respPendingMatch.TerminalPendingMatchs + respPendingMatch.AccesoriesPendingMatchs
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cacOrderTerminalRequest := modeladapter.CacOrderTerminalsRequest{Dealer: &user.EntitiesInfo.Dealer, UsrPlace: &user.WarehouseInfo.UserPlace}
		resp, ok = adapter.CacOrderTerminals(cacOrderTerminalRequest)
		if !ok {
			logging.Info("::: Error in adapter cacOrderTerminals :::")
			return
		}
		if resp.StatusCode != 200 {
			return
		}

		cacOrderTerminals, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logging.Error("::: Error to read cacOrderTerminals:::")
			return
		}
		type resposeCacOrderTerminals []struct {
			AlmacenSap      string `json:"almacenSap"`
			Dealer          string `json:"dealer"`
			Entidad         string `json:"entidad"`
			ID              string `json:"id"`
			LastUpdateDate  string `json:"lastUpdateDate"`
			Pin             string `json:"pin"`
			ProductSerial   string `json:"productSerial"`
			Producto        string `json:"producto"`
			ReasonRejection string `json:"reasonRejection"`
			Reserva         string `json:"reserva"`
			Status          string `json:"status"`
			UserPlace       string `json:"userPlace"`
		}
		respCacOrderTerminals := resposeCacOrderTerminals{}
		json.Unmarshal(cacOrderTerminals, &respCacOrderTerminals)

		user.AccountingItemsInfo.PendingIntermidiateMaterials = 0
		user.AccountingItemsInfo.PendingProducts = 0
		for _, element := range respCacOrderTerminals {
			if element.Status == INTERMEDIATE {
				user.AccountingItemsInfo.PendingIntermidiateMaterials++
			}
			if element.Status == ENTEGADA || element.Status == ESTADOH {
				user.AccountingItemsInfo.PendingProducts++
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var usrIds []string
		usrIds = append(usrIds, usrId)
		userTaskRequest := modeladapter.UserTaskRequest{UsrIds: usrIds}
		resp, ok = adapter.UserTask(userTaskRequest)
		if !ok {
			logging.Info("::: Error in adapter userTask :::")
			return
		}
		if resp.StatusCode != 200 {
			return
		}
		userTask, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logging.Error("::: Error to read userTask :::")
			return
		}
		type responseUserTask []struct {
			Tasks []string `json:"tasks"`
			User  string   `json:"user"`
		}
		respUserTask := responseUserTask{}
		json.Unmarshal(userTask, &respUserTask)
		user.UserInfo.Tasks = respUserTask[0].Tasks
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var usrIds []string
		usrIds = append(usrIds, usrId)
		userTaskPriorityRequest := modeladapter.UserTaskPriorityRequest{UsrIds: usrIds}
		resp, ok = adapter.UserTaskPriority(userTaskPriorityRequest)
		if !ok {
			logging.Info("::: Error in adapter userTaskPriority :::")
			return
		}
		if resp.StatusCode != 200 {
			return
		}
		userTaskPriority, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logging.Error("::: Error to read userTaskPriority :::")
			return
		}
		type responseUserTaskPriority []struct {
			User         string `json:"user"`
			TaskPriority string `json:"taskPriority"`
		}
		respTaskPriority := responseUserTaskPriority{}
		json.Unmarshal(userTaskPriority, &respTaskPriority)
		user.UserInfo.TaskPriority = respTaskPriority[0].TaskPriority
	}()

	wg.Wait()
	context.JSON(http.StatusOK, user)
}
