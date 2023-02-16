package service

import (
	"github.com/danjvs/go-CRUD/src/configuration/rest_err"
	"github.com/danjvs/go-CRUD/src/model"
	"github.com/danjvs/go-CRUD/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)
}
