package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func vaidatorRequest(i interface{}) string {
	//can add specific tagg names//
	if err := validate.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors.Error()
	}
	return ""
}

func newError(ctx *gin.Context, status int, err string) {
	er := HTTPError{
		Code:    status,
		Message: err,
	}
	ctx.JSON(status, er)
}

func stringIsNullOrEmpty(s string) bool {
	return len(s) < 1 || s == ""
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
