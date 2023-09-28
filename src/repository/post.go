package repository

import (
	"echo-pet-api/database"
	"echo-pet-api/src/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository() *PostRepository {
	return &PostRepository{db: database.Connection()}
}

func (pr *PostRepository) GetAll() ([]*model.Post, error) {
	var posts []*model.Post
	err := pr.db.Find(&posts).Error

	return posts, err
}

func (pr *PostRepository) GetById(id uint) (*model.Post, error) {
	var post *model.Post
	err := pr.db.First(&post, id).Error

	return post, err
}

func (pr *PostRepository) GetBySlug(slug string) (*model.Post, error) {
	var post *model.Post
	err := pr.db.Where("slug = ?", slug).First(&post).Error

	return post, err
}

func (pr *PostRepository) Create(post *model.Post) error {
	return pr.db.Create(post).Error
}

func (pr *PostRepository) Update(post *model.Post) error {
	return pr.db.Save(post).Error
}

func (pr *PostRepository) Delete(post *model.Post) error {
	return pr.db.Delete(post).Error
}
