package dto

import (
	"blog-service/src/model"
)

type UserStoreRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
}

func NewUserResponse(user *model.User) *UserResponse {
	response := &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Unix(),
	}

	return response
}

type ShortUserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewShortUserResponse(user *model.User) *ShortUserResponse {
	return &ShortUserResponse{Email: user.Email, Name: user.Name}
}
