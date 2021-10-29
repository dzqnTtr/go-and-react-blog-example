package service

import (
	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type IPostRepository interface {
	GetAllPost() ([]model.Post, error)
}

type PostService struct {
	repository IPostRepository
}

func NewPostService(repository IPostRepository) PostService {
	return PostService{
		repository: repository,
	}
}

func (s PostService) GetAllPost() ([]model.Post, error) {
	return s.repository.GetAllPost()
}
