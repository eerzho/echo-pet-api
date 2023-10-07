package service

import (
	"echo-pet-api/src/model"
	"echo-pet-api/src/model/dto"
	"echo-pet-api/src/repository"
)

type UserService struct {
	repository *repository.UserRepository
	jwt        *JWTService
}

func NewUserService() *UserService {
	return &UserService{
		repository: repository.NewUserRepository(),
		jwt:        NewJWTService(),
	}
}

func (us *UserService) GetAll() ([]*dto.UserResponse, error) {
	users, err := us.repository.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]*dto.UserResponse, len(users))
	for index, user := range users {
		response[index] = dto.NewUserResponse(user)
	}

	return response, nil
}

func (us *UserService) GetById(id uint) (*dto.UserResponse, error) {
	user, err := us.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewUserResponse(user), nil
}

func (us *UserService) Create(request *dto.UserStoreRequest) (*dto.UserResponse, error) {
	passwordHash, err := us.jwt.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{Email: request.Email, Name: request.Name, Password: passwordHash}

	err = us.repository.Create(&user)
	if err != nil {
		return nil, err
	}

	return dto.NewUserResponse(&user), nil
}

func (us *UserService) UpdatePassword(id uint, request *dto.UserUpdateRequest) (*dto.UserResponse, error) {
	user, err := us.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	user.Password, err = us.jwt.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	if err = us.repository.UpdatePassword(user); err != nil {
		return nil, err
	}

	return dto.NewUserResponse(user), nil
}

func (us *UserService) Delete(id uint) error {
	user, err := us.repository.GetById(id)
	if err != nil {
		return err
	}

	err = us.repository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
