package model

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Slug  string `gorm:"unique_index,not null"`
	Title string `gorm:"not null"`
	Desc  string
}

type PostStoreRequest struct {
	Slug  string `json:"slug"`
	Title string `json:"title" validate:"required"`
	Desc  string `json:"desc"`
}

type PostUpdateRequest struct {
	Desc string `json:"desc" validate:"required"`
}

type PostResponse struct {
	ID        uint   `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	CreatedAt int64  `json:"createdAt"`
}

func NewPostFromStoreRequest(request *PostStoreRequest) *Post {
	if request.Slug == "" {
		request.Slug = slug.Make(request.Title)
	}

	return &Post{Title: request.Title, Slug: request.Slug, Desc: request.Desc}
}

func NewPostResponseFromModel(post *Post) *PostResponse {
	return &PostResponse{
		ID:        post.ID,
		Slug:      post.Slug,
		Title:     post.Title,
		Desc:      post.Desc,
		CreatedAt: post.CreatedAt.Unix(),
	}
}
