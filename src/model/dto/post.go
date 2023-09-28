package dto

import "echo-pet-api/src/model"

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

func NewPostResponse(post *model.Post) *PostResponse {
	return &PostResponse{
		ID:        post.ID,
		Slug:      post.Slug,
		Title:     post.Title,
		Desc:      post.Desc,
		CreatedAt: post.CreatedAt.Unix(),
	}
}
