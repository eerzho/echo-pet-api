package service

import (
	"auth-service/src/exception"
	"auth-service/src/model/dto"
	"auth-service/src/service/service_interface"
	"errors"
	"gorm.io/gorm"
)

type AuthService struct {
	userService service_interface.UserServiceInterface
	jwtService  service_interface.JWTServiceInterface
}

func NewAuthService() *AuthService {
	return &AuthService{
		userService: NewUserService(),
		jwtService:  NewJWTService(),
	}
}

func (this *AuthService) Login(request *dto.LoginRequest) (string, error) {
	user, err := this.userService.GetByEmail(request.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", exception.ErrInvalidLogin
	} else if err != nil {
		return "", err
	}

	if !this.jwtService.IsEqual(user.Password, request.Password) {
		return "", exception.ErrInvalidLogin
	}

	token, err := this.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
