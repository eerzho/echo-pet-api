package repository_i

import "auth-service/src/model"

type UserRepositoryI interface {
	GetAll() ([]*model.User, error)
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Save(user model.User) (*model.User, error)
	Delete(id uint) error
	HasPermission(id uint, permissionSlug string) bool
}
