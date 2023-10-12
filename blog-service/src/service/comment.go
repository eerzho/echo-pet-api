package service

import (
	"blog-service/src/exception"
	"blog-service/src/model"
	"blog-service/src/model/dto"
	"blog-service/src/repository"
)

type CommentService struct {
	repository *repository.CommentRepository
}

func NewCommentService() *CommentService {
	return &CommentService{repository: repository.NewCommentRepository()}
}

func (cs *CommentService) GetAll() ([]*dto.CommentResponse, error) {
	comments, err := cs.repository.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]*dto.CommentResponse, len(comments))
	for index, comment := range comments {
		response[index] = dto.NewCommentResponse(comment)
	}

	return response, nil
}

func (cs *CommentService) GetById(id uint) (*dto.CommentResponse, error) {
	comment, err := cs.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewCommentResponse(comment), nil
}

func (cs *CommentService) Create(userID uint, request *dto.CommentStoreRequest) (*dto.CommentResponse, error) {
	comment := model.Comment{Text: request.Text, UserID: userID, PostID: request.PostID}

	err := cs.repository.Create(&comment)
	if err != nil {
		return nil, err
	}

	return dto.NewCommentResponse(&comment), nil
}

func (cs *CommentService) Update(userID, id uint, request *dto.CommentUpdateRequest) (*dto.CommentResponse, error) {
	comment, err := cs.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if comment.UserID != userID {
		return nil, exception.NewPermissionDenied()
	}

	comment.Text = request.Text
	if err = cs.repository.Update(comment); err != nil {
		return nil, err
	}

	return dto.NewCommentResponse(comment), nil
}

func (cs *CommentService) Delete(userID, id uint) error {
	comment, err := cs.repository.GetById(id)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return exception.NewPermissionDenied()
	}

	if err = cs.repository.Delete(comment); err != nil {
		return err
	}

	return nil
}
