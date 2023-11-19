package repository_i

import "auth-service/src/model"

type PermissionRepositoryI interface {
	GetAllByRole(roleID uint) ([]*model.Permission, error)
	GetById(id uint) (*model.Permission, error)
	Save(permission model.Permission) (*model.Permission, error)
	Delete(id uint) error
}
