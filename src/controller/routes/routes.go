package routes

import (
	controller "github.com/danjvs/go-CRUD/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup) {

	route.GET("/user/:userId", controller.FindUserById)
	route.GET("/user/mail/:userMail", controller.FindUserByEmail)
	route.POST("/user", controller.CreateUser)
	route.PUT("/user/:userId", controller.UpdateUser)
	route.DELETE("/user/:userId", controller.DeleteUser)
}
