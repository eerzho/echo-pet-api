package service_i

import (
	"auth-service/src/dto"
	"auth-service/src/model"
)

type RoleServiceI interface {
	GetAll() ([]*model.Role, error)
	GetById(id uint) (*model.Role, error)
	GetBySlug(slug string) (*model.Role, error)
	Create(request *dto.RoleStoreRequest) (*model.Role, error)
	Update(id uint, request *dto.RoleUpdateRequest) (*model.Role, error)
	Delete(id uint) error
	AddPermissions(id uint, request *dto.RoleAddPermissionsRequest) error
	RemovePermissions(id uint, request *dto.RoleRemovePermissionsRequest) error
}
