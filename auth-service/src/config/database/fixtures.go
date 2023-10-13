package database

import (
	"auth-service/src/model"
	"errors"
	"gorm.io/gorm"
)

func RunFixtures(db *gorm.DB) {
	runRoles(db)
	runPermissions(db)
	assignPermissionsToAdmin(db)
}

func runRoles(db *gorm.DB) {
	var roles = []model.Role{
		{Name: "Admin", Slug: "admin"},
		{Name: "User", Slug: "user"},
	}

	for _, role := range roles {
		var existingRole model.Role
		if err := db.Where("slug = ?", role.Slug).First(&existingRole).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&role)
			}
		}
	}
}

func runPermissions(db *gorm.DB) {
	var permissions = []model.Permission{
		{Name: "Users index", Slug: "users-index"},
		{Name: "Roles index", Slug: "roles-index"},
		{Name: "Roles store", Slug: "roles-store"},
		{Name: "Roles show", Slug: "roles-show"},
		{Name: "Roles update", Slug: "roles-update"},
		{Name: "Roles delete", Slug: "roles-delete"},
		{Name: "Roles add permissions", Slug: "roles-add-permissions"},
		{Name: "Roles remove permissions", Slug: "roles-remove-permissions"},
		{Name: "Permissions index", Slug: "permissions-index"},
		{Name: "Permissions store", Slug: "permissions-store"},
		{Name: "Permissions show", Slug: "permissions-show"},
		{Name: "Permissions update", Slug: "permissions-update"},
		{Name: "Permissions delete", Slug: "permissions-delete"},
	}

	for _, permission := range permissions {
		var existingPermission model.Permission
		if err := db.Where("slug = ?", permission.Slug).First(&existingPermission).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&permission)
			}
		}
	}
}

func assignPermissionsToAdmin(db *gorm.DB) {
	var adminRole model.Role
	if err := db.Where("slug = ?", "admin").First(&adminRole).Error; err != nil {
		return
	}

	var permissions []model.Permission
	if err := db.Find(&permissions).Error; err != nil {
		return
	}

	if err := db.Model(&adminRole).Association("Permissions").Append(&permissions); err != nil {
		return
	}
}
