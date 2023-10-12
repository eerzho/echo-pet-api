package repository

import (
	"auth-service/src/config/database"
	"auth-service/src/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{connection: database.Connection()}
}

func (this *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := this.connection.Find(&users).Error

	return users, err
}

func (this *UserRepository) GetById(id uint) (*model.User, error) {
	var user *model.User
	err := this.connection.First(&user, id).Error

	return user, err
}

func (this *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user *model.User
	err := this.connection.Where("email = ?", email).First(&user).Error

	return user, err
}

func (this *UserRepository) Save(user model.User) (*model.User, error) {
	result := this.connection.Save(&user)

	return &user, result.Error
}

func (this *UserRepository) Delete(id uint) error {
	return this.connection.Delete(&model.User{}, id).Error
}
