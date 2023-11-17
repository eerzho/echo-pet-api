package repository

import (
	"auth-service/src/model"
)

type PermissionRepository struct {
	*BaseRepository[model.Permission]
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{BaseRepository: NewBaseRepository[model.Permission]()}
}

func (this *PermissionRepository) GetAllByRole(roleID uint) ([]*model.Permission, error) {
	var role model.Role

	err := this.connection.Preload("Permissions").First(&role, roleID).Error
	if err != nil {
		return nil, err
	}

	return role.Permissions, nil
}
