package service_interface

import "auth-service/src/model/dto"

type AuthServiceInterface interface {
	Login(request *dto.LoginRequest) (string, error)
}
