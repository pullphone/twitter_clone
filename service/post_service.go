package service

import (
	"github.com/pullphone/twitter_clone/entity"
	"github.com/pullphone/twitter_clone/repository"
)

type PostService struct {
	PostRepository repository.PostRepository
}

func (service *PostService) Add(post entity.Post) (id int64, err error) {
	id, err = service.PostRepository.Store(post)
	return
}

func (service *PostService) PostById(id int64) (post entity.Post, err error) {
	post, err = service.PostRepository.FindById(id)
	return
}

func (service *PostService) Posts() (posts []entity.Post, err error) {
	posts, err = service.PostRepository.FindAll()
	return
}
