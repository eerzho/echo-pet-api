package service_interface

import (
	"auth-service/src/model"
	"auth-service/src/model/dto"
)

type RoleServiceInterface interface {
	GetAll() ([]*model.Role, error)
	GetById(id uint) (*model.Role, error)
	GetBySlug(slug string) (*model.Role, error)
	Create(request *dto.RoleStoreRequest) (*model.Role, error)
	Update(id uint, request *dto.RoleUpdateRequest) (*model.Role, error)
	Delete(id uint) error
	AddPermissions(id uint, request *dto.RoleAddPermissionsRequest) error
	RemovePermissions(id uint, request *dto.RoleRemovePermissionsRequest) error
}
