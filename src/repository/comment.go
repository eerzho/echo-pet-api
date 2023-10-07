package repository

import (
	"echo-pet-api/src/database"
	"echo-pet-api/src/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{db: database.Connection()}
}

func (cr *CommentRepository) GetAll() ([]*model.Comment, error) {
	var comments []*model.Comment
	err := cr.db.Preload("User").Find(&comments).Error

	return comments, err
}

func (cr *CommentRepository) GetById(id uint) (*model.Comment, error) {
	var comment *model.Comment
	err := cr.db.Preload("User").First(&comment, id).Error

	return comment, err
}

func (cr *CommentRepository) Create(comment *model.Comment) error {
	if err := cr.db.Save(comment).Error; err != nil {
		return err
	}
	return cr.db.Preload("User").First(&comment, comment.ID).Error
}

func (cr *CommentRepository) Update(comment *model.Comment) error {
	return cr.db.Save(comment).Error
}

func (cr *CommentRepository) Delete(comment *model.Comment) error {
	return cr.db.Delete(comment).Error
}
