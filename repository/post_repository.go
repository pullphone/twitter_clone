package repository

import (
	"github.com/pullphone/twitter_clone/entity"
	"time"
	"database/sql"
)

type PostRepository struct {
	SqlHandler
}

func (repo *PostRepository) Store(p entity.Post) (id int64, err error) {
	result, err := repo.Execute("INSERT INTO posts (text, created_at) VALUES (?, CURRENT_TIMESTAMP)", p.Text)
	if err != nil {
		return
	}

	id, err = result.LastInsertId()
	return
}

func (repo *PostRepository) FindById(id int64) (post entity.Post, err error) {
	rows, err := repo.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		return
	}

	var id2 int64
	var text string
	var utime time.Time
	var ctime time.Time
	rows.Next()
	if err = rows.Scan(&id2, &text, &utime, &ctime); err != nil {
		return post, sql.ErrNoRows
	}

	post.ID = id2
	post.Text = text
	post.UpdatedAt = utime.Unix()
	post.CreatedAt = ctime.Unix()
	return
}
