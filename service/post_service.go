package service

import (
	"github.com/pullphone/twitter_clone/dao"
	"github.com/pullphone/twitter_clone/entity"
	"github.com/pullphone/twitter_clone/repository"
)

type PostService interface {
	Add(entity.Post) (int64, error)
	GetById(int64) (entity.Post, error)
	GetAll() ([]entity.Post, error)
}

type postService struct {
	PostRepository repository.PostRepository
}

func NewPostService() *postService {
	return &postService{
		PostRepository: repository.PostRepository{
			dao.NewPostDao(),
		},
	}
}

func (service *postService) Add(post entity.Post) (id int64, err error) {
	id, err = service.PostRepository.Store(post)
	return
}

func (service *postService) GetById(id int64) (post entity.Post, err error) {
	post, err = service.PostRepository.FindById(id)
	return
}

func (service *postService) GetAll() (posts []entity.Post, err error) {
	posts, err = service.PostRepository.FindAll()
	return
}
