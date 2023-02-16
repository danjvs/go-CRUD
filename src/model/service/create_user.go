package service

import (
	"github.com/danjvs/go-CRUD/src/configuration/logger"
	"github.com/danjvs/go-CRUD/src/configuration/rest_err"
	"github.com/danjvs/go-CRUD/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
