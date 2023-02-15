package controller

import (
	"fmt"
	"net/http"

	"github.com/danjvs/go-CRUD/src/configuration/logger"
	"github.com/danjvs/go-CRUD/src/configuration/validation"
	"github.com/danjvs/go-CRUD/src/controller/model/request"
	"github.com/danjvs/go-CRUD/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FindUserById(context *gin.Context) {
}

func FindUserByEmail(context *gin.Context) {
}

func CreateUser(context *gin.Context) {
	logger.Info("Init create user controller", zap.String("Journey", "create user"))
	var UserRequest request.UserRequest

	if err := context.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("Journey", "create user"))
		errRest := validation.ValidateUserError(err)

		context.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(UserRequest)
	response := response.UserResponse{
		ID:    "01",
		Email: UserRequest.Email,
		Name:  UserRequest.Name,
		Age:   UserRequest.Age,
	}

	logger.Info("User created successfully", zap.String("Journey", "create user"))
	context.JSON(http.StatusOK, response)

}

func DeleteUser(context *gin.Context) {
}

func UpdateUser(context *gin.Context) {
}
