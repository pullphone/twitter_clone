package repository

import (
	"github.com/pullphone/twitter_clone/entity"
	"github.com/pullphone/twitter_clone/dao"
	"strconv"
)

type PostRepository struct {
	PostDao dao.PostDao
}

func (repo *PostRepository) Store(p entity.Post) (id int64, err error) {
	id, err = repo.PostDao.Insert(p.Text)
	return
}

func (repo *PostRepository) FindById(id int64) (post entity.Post, err error) {
	row, err := repo.PostDao.Get(id)
	if err != nil {
		return
	}

	post = parseEntity(row)
	return
}

func (repo *PostRepository) FindAll() (posts []entity.Post, err error) {
	rows, err := repo.PostDao.GetList()
	if err != nil {
		return
	}

	for _, row := range rows {
		post := parseEntity(row)
		posts = append(posts, post)
	}
	return
}

func parseEntity(row dao.Post) (entity.Post) {
	return entity.Post{
		ID: strconv.FormatInt(row.ID, 10),
		Text: row.Text,
		UpdatedAt: row.UpdatedAt.Unix(),
		CreatedAt: row.CreatedAt.Unix(),
	}
}
