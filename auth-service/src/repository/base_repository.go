package repository

import (
	"auth-service/src/application"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	connection *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{connection: application.GlobalDB}
}

func (this *BaseRepository[T]) GetById(id uint) (*T, error) {
	var model *T
	err := this.connection.First(&model, id).Error

	return model, err
}

func (this *BaseRepository[T]) Save(model T) (*T, error) {
	result := this.connection.Save(&model)

	return &model, result.Error
}

func (this *BaseRepository[T]) Delete(id uint) error {
	var model T
	if err := this.connection.First(&model, id).Error; err != nil {
		return err
	}

	return this.connection.Delete(&model).Error
}
