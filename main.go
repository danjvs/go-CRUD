package main

import (
	"context"
	"log"

	"github.com/danjvs/go-CRUD/src/configuration/database/mongodb"
	"github.com/danjvs/go-CRUD/src/configuration/logger"
	"github.com/danjvs/go-CRUD/src/controller"
	"github.com/danjvs/go-CRUD/src/controller/routes"
	"github.com/danjvs/go-CRUD/src/model/repository"
	"github.com/danjvs/go-CRUD/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	//userController := initDependencies(database)
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
