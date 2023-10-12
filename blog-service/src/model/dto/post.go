package dto

import (
	"blog-service/src/model"
)

type PostStoreRequest struct {
	Slug  string `json:"slug"`
	Title string `json:"title" validate:"required"`
	Desc  string `json:"desc"`
}

type PostUpdateRequest struct {
	Desc string `json:"desc" validate:"required"`
}

type PostResponse struct {
	ID        uint               `json:"id"`
	Slug      string             `json:"slug"`
	Title     string             `json:"title"`
	Desc      string             `json:"desc"`
	CreatedAt int64              `json:"created_at"`
	AuthorID  uint               `json:"author_id"`
	Author    *ShortUserResponse `json:"author"`
}

func NewPostResponse(post *model.Post) *PostResponse {
	response := &PostResponse{
		ID:        post.ID,
		Slug:      post.Slug,
		Title:     post.Title,
		Desc:      post.Desc,
		CreatedAt: post.CreatedAt.Unix(),
		AuthorID:  post.AuthorID,
	}

	if post.Author.ID != 0 {
		response.Author = NewShortUserResponse(&post.Author)
	}

	return response
}
