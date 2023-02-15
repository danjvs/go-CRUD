package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danjvs/go-CRUD/src/configuration/validation"
	"github.com/danjvs/go-CRUD/src/controller/model/request"
	"github.com/danjvs/go-CRUD/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func FindUserById(context *gin.Context) {
}

func FindUserByEmail(context *gin.Context) {
}

func CreateUser(context *gin.Context) {
	log.Println("Init create user controller")
	var UserRequest request.UserRequest

	if err := context.ShouldBindJSON(&UserRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
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

	context.JSON(http.StatusOK, response)

}

func DeleteUser(context *gin.Context) {
}

func UpdateUser(context *gin.Context) {
}
