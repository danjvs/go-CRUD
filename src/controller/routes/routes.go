package routes

import (
	controller "github.com/danjvs/go-CRUD/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup) {

	route.GET("/usuario/:userId", controller.FindUserById)
	route.GET("/usuario/email/:userEmail", controller.FindUserByEmail)
	route.POST("/usuario", controller.CreateUser)
	route.PUT("/usuario/:userId", controller.UpdateUser)
	route.DELETE("/usuario/:userId", controller.DeleteUser)
}
