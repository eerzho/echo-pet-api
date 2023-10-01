package dto

import "echo-pet-api/src/model"

type UserStoreRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint                `json:"id"`
	Email     string              `json:"email"`
	Name      string              `json:"name"`
	CreatedAt int64               `json:"createdAt"`
	Posts     []*UserPostResponse `json:"posts"`
}

func NewUserResponse(user *model.User) *UserResponse {
	response := &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Unix(),
	}

	var posts []*UserPostResponse
	for _, post := range user.Posts {
		posts = append(posts, NewUserPostResponse(&post))
	}

	response.Posts = posts

	return response
}

type PostAuthorResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func NewPostAuthorResponse(user *model.User) *PostAuthorResponse {
	return &PostAuthorResponse{ID: user.ID, Email: user.Email}
}
