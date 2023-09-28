package service

import (
	"echo-pet-api/src/exception"
	"echo-pet-api/src/model/dto"
	"echo-pet-api/src/repository"
	"errors"
	"gorm.io/gorm"
)

type AuthService struct {
	repository *repository.UserRepository
	jwt        *JWTService
}

func NewAuthService() *AuthService {
	return &AuthService{
		repository: repository.NewUserRepository(),
		jwt:        NewJWTService(),
	}
}

func (as *AuthService) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := as.repository.GetByEmail(request.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.NewInvalidLoginError()
	} else if err != nil {
		return nil, err
	}

	if !as.jwt.IsEqual(user.Password, request.Password) {
		return nil, exception.NewInvalidLoginError()
	}

	token, err := as.jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return dto.NewLoginResponse(token), nil
}
