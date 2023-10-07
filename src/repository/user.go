package repository

import (
	"echo-pet-api/src/database"
	"echo-pet-api/src/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.Connection()}
}

func (ur *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := ur.db.Find(&users).Error

	return users, err
}

func (ur *UserRepository) GetById(id uint) (*model.User, error) {
	var user *model.User
	err := ur.db.First(&user, id).Error

	return user, err
}

func (ur *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user *model.User
	err := ur.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (ur *UserRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) UpdatePassword(user *model.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(user *model.User) error {
	return ur.db.Delete(&user).Error
}
