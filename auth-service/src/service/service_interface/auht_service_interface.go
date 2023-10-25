package service_interface

import "auth-service/src/dto"

type AuthServiceInterface interface {
	Login(request *dto.LoginRequest) (string, error)
}
