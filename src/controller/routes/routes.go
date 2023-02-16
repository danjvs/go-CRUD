package routes

import (
	controller "github.com/danjvs/go-CRUD/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup, userController controller.UserControllerInterface) {

	route.POST("/user", userController.CreateUser)
}
