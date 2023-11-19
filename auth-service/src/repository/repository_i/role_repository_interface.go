package repository_i

import "auth-service/src/model"

type RoleRepositoryI interface {
	GetAll() ([]*model.Role, error)
	GetById(id uint) (*model.Role, error)
	GetBySlug(slug string) (*model.Role, error)
	Save(role model.Role) (*model.Role, error)
	Delete(id uint) error
	AddPermissions(id uint, permissionsID []uint) error
	RemovePermissions(id uint, permissionsID []uint) error
}
