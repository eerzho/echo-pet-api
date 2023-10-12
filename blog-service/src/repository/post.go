package repository

import (
	"blog-service/src/database"
	"blog-service/src/model"
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
	err := pr.db.Preload("Author").Find(&posts).Error

	return posts, err
}

func (pr *PostRepository) GetById(id uint) (*model.Post, error) {
	var post *model.Post
	err := pr.db.Preload("Author").First(&post, id).Error

	return post, err
}

func (pr *PostRepository) GetBySlug(slug string) (*model.Post, error) {
	var post *model.Post
	err := pr.db.Preload("Author").Where("slug = ?", slug).First(&post).Error

	return post, err
}

func (pr *PostRepository) GetBySlugCount(slug string) (int, error) {
	var count int64
	err := pr.db.Table("posts").Where("slug LIKE ?", slug+"%").Count(&count).Error

	return int(count), err
}

func (pr *PostRepository) Create(post *model.Post) error {
	if err := pr.db.Save(post).Error; err != nil {
		return err
	}

	return pr.db.Preload("Author").First(&post, post.ID).Error
}

func (pr *PostRepository) Update(post *model.Post) error {
	return pr.db.Save(post).Error
}

func (pr *PostRepository) Delete(post *model.Post) error {
	return pr.db.Delete(post).Error
}
