package service

import (
	"auth-service/src/dto"
	"auth-service/src/exception"
	"auth-service/src/model"
	"auth-service/src/service/service_i"
	"errors"
	"gorm.io/gorm"
)

type AuthService struct {
	userService service_i.UserServiceI
	jwtService  service_i.JWTServiceI
}

func NewAuthService(userService service_i.UserServiceI, jwtService service_i.JWTServiceI) *AuthService {
	return &AuthService{
		userService: userService,
		jwtService:  jwtService,
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

func (this *AuthService) GetUserById(id uint) (*model.User, error) {
	return this.userService.GetById(id)
}
