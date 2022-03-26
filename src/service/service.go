package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	adapter "demoproject/src/adapter"
	model "demoproject/src/model"

	"github.com/gin-gonic/gin"
)

const ERROR_DESCRIPTION = "Error description."
const SUCCESS_DESCRIPTION = "Success description."
const ERROR_STATUS = 1
const SUCCESS_STATUS = 0

// user godoc
// @Summary Show user
// @Description Obtains info of Users
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body model.UserInsert true "User"
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Success 200 {object} model.Status
// @Router /user [post]
func PostUser(context *gin.Context) {
	var userInsert model.UserInsert
	json.NewDecoder(context.Request.Body).Decode(&userInsert)
	err := vaidatorRequest(userInsert)
	if !stringIsNullOrEmpty(err) {
		newError(context, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	statusModel := model.Status{}
	statusModel.Description = SUCCESS_DESCRIPTION
	statusModel.Status = SUCCESS_STATUS
	context.JSON(http.StatusOK, statusModel)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	user := model.User{}
	user.Email = "joaquin.demaria@mail.com"
	user.FirstName = "Joaquin"
	user.LastName = "Demaria"
	user.ID = 1
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func ApiIntegrationGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	response := adapter.ApiIntegrationGet()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
