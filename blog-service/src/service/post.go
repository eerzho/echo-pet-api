package service

import (
	"blog-service/src/exception"
	"blog-service/src/model"
	"blog-service/src/model/dto"
	"blog-service/src/repository"
	"fmt"
	"github.com/gosimple/slug"
)

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService() *PostService {
	return &PostService{repository: repository.NewPostRepository()}
}

func (ps *PostService) GetAll() ([]*dto.PostResponse, error) {
	posts, err := ps.repository.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]*dto.PostResponse, len(posts))
	for index, post := range posts {
		response[index] = dto.NewPostResponse(post)
	}

	return response, nil
}

func (ps *PostService) GetById(id uint) (*dto.PostResponse, error) {
	post, err := ps.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewPostResponse(post), nil
}

func (ps *PostService) GetBySlug(slug string) (*dto.PostResponse, error) {
	post, err := ps.repository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}

	return dto.NewPostResponse(post), nil
}

func (ps *PostService) Create(authorId uint, request *dto.PostStoreRequest) (*dto.PostResponse, error) {

	if request.Slug == "" {
		request.Slug = slug.Make(request.Title)

		slugCount, err := ps.repository.GetBySlugCount(request.Slug)
		if err != nil {
			return nil, err
		}

		if slugCount > 0 {
			request.Slug = fmt.Sprintf("%s-%d", request.Slug, slugCount)
		}
	}

	post := model.Post{Slug: request.Slug, Title: request.Title, Desc: request.Desc, AuthorID: authorId}

	err := ps.repository.Create(&post)
	if err != nil {
		return nil, err
	}

	return dto.NewPostResponse(&post), nil
}

func (ps *PostService) Update(authorID, id uint, request *dto.PostUpdateRequest) (*dto.PostResponse, error) {
	post, err := ps.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if post.AuthorID != authorID {
		return nil, exception.NewPermissionDenied()
	}

	post.Desc = request.Desc
	if err = ps.repository.Update(post); err != nil {
		return nil, err
	}

	return dto.NewPostResponse(post), nil
}

func (ps *PostService) Delete(authorID, id uint) error {
	post, err := ps.repository.GetById(id)
	if err != nil {
		return err
	}

	if post.AuthorID != authorID {
		return exception.NewPermissionDenied()
	}

	err = ps.repository.Delete(post)
	if err != nil {
		return err
	}

	return nil
}
