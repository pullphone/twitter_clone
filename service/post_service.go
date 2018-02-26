package service

import "github.com/pullphone/twitter_clone/entity"

type PostService struct {
	PostRepository
}

func (service *PostService) Add(post entity.Post) (id int64, err error) {
	id, err = service.PostRepository.Store(post)
	return
}

func (service *PostService) PostById(id int64) (post entity.Post, err error) {
	post, err = service.PostRepository.FindById(id)
	return
}

type PostRepository interface {
	Store(entity.Post) (int64, error)
	FindById(int64) (entity.Post, error)
}
