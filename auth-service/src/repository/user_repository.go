package repository

import (
	"auth-service/src/model"
)

type UserRepository struct {
	*BaseRepository[model.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{BaseRepository: NewBaseRepository[model.User]()}
}

func (this *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := this.connection.Find(&users).Error

	return users, err
}

func (this *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user *model.User
	err := this.connection.Where("email = ?", email).First(&user).Error

	return user, err
}

func (this *UserRepository) HasPermission(id uint, permissionSlug string) bool {
	var count int64
	err := this.connection.
		Model(&model.Permission{}).
		Joins("JOIN roles_permissions ON permissions.id = roles_permissions.permission_id").
		Joins("JOIN roles ON roles.id = roles_permissions.role_id").
		Joins("JOIN users ON users.role_id = roles.id").
		Where("users.id = ? AND permissions.slug = ?", id, permissionSlug).
		Count(&count).Error

	if err != nil {
		return false
	}

	return count > 0
}
