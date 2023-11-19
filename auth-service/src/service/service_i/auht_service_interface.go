package service_i

import (
	"auth-service/src/dto"
	"auth-service/src/model"
)

type AuthServiceI interface {
	Login(request *dto.LoginRequest) (string, error)
	GetUserById(id uint) (*model.User, error)
}
