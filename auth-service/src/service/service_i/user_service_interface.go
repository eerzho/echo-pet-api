package service_i

import (
	"auth-service/src/dto"
	"auth-service/src/model"
)

type UserServiceI interface {
	GetAll() ([]*model.User, error)
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(request *dto.UserStoreRequest) (*model.User, error)
	UpdatePassword(id uint, request *dto.UserUpdatePasswordRequest) (*model.User, error)
	Delete(id uint) error
	HasPermission(id uint, permissionSlug string) bool
}
