package repository

import (
	"auth-service/src/config/database"
	"auth-service/src/model"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	connection *gorm.DB
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{connection: database.Connection()}
}

func (this *PermissionRepository) GetAllByRole(roleID uint) ([]*model.Permission, error) {
	var role model.Role

	err := this.connection.Preload("Permissions").First(&role, roleID).Error
	if err != nil {
		return nil, err
	}

	return role.Permissions, nil
}

func (this *PermissionRepository) GetById(id uint) (*model.Permission, error) {
	var permission *model.Permission
	err := this.connection.First(&permission, id).Error

	return permission, err
}

func (this *PermissionRepository) Save(permission model.Permission) (*model.Permission, error) {
	result := this.connection.Save(&permission)

	return &permission, result.Error
}

func (this *PermissionRepository) Delete(id uint) error {
	return this.connection.Delete(&model.Permission{}, id).Error
}
