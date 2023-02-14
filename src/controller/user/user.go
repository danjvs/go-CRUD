package controller

import (
	"fmt"

	rest_err "github.com/danjvs/go-CRUD/src/configuration"
	"github.com/danjvs/go-CRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func FindUserById(context *gin.Context) {
}

func FindUserByEmail(context *gin.Context) {
}

func CreateUser(context *gin.Context) {

	var UserRequest request.UserRequest

	if err := context.ShouldBindJSON(&UserRequest); err != nil {

		restErr := rest_err.NewBadRequestError((fmt.Sprintf("There are some incorrect filds, error=%s", err.Error())))

		context.JSON(restErr.Code, restErr)
		return
	}
}

func DeleteUser(context *gin.Context) {
}

func UpdateUser(context *gin.Context) {
}
