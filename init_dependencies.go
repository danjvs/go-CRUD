package main

import (
	"github.com/danjvs/go-CRUD/src/controller"
	"github.com/danjvs/go-CRUD/src/model/repository"
	"github.com/danjvs/go-CRUD/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
