package dto

import "echo-pet-api/src/model"

type CommentStoreRequest struct {
	Text   string `json:"text" validate:"required,min=10,max=500"`
	PostID uint   `json:"post_id" validate:"required,gt=0"`
}

type CommentUpdateRequest struct {
	Text string `json:"text" validate:"required,min=10,max=500"`
}

type CommentResponse struct {
	ID        uint               `json:"id"`
	Text      string             `json:"text"`
	CreatedAt int64              `json:"created_at"`
	PostID    uint               `json:"post_id"`
	UserID    uint               `json:"user_id"`
	User      *ShortUserResponse `json:"user"`
}

func NewCommentResponse(comment *model.Comment) *CommentResponse {
	response := &CommentResponse{
		ID:        comment.ID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt.Unix(),
		PostID:    comment.PostID,
		UserID:    comment.UserID,
	}

	if comment.User.ID != 0 {
		response.User = NewShortUserResponse(&comment.User)
	}

	return response
}
