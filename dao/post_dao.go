package dao

import (
	"github.com/pullphone/twitter_clone/infrastructure"
	"time"
	"database/sql"
)

type PostDao struct {
}

func (dao *PostDao) Insert(text string) (id int64, err error) {
	result, err := infrastructure.DBExecute("INSERT INTO posts (text, created_at) VALUES (?, CURRENT_TIMESTAMP)", text)
	if err != nil {
		return
	}

	id, err = result.LastInsertId()
	return
}

func (dao *PostDao) Get(targetId int64) (post Post, err error) {
	rows, err := infrastructure.DBQuery("SELECT * FROM posts WHERE id = ?", targetId)
	defer rows.Close()
	if err != nil {
		return post, err
	}

	result := rows.Next()
	if !result {
		return post, sql.ErrNoRows
	}
	err = rows.StructScan(&post)

	return
}

func (dao *PostDao) GetList() (posts []Post, err error) {
	rows, err := infrastructure.DBQuery("SELECT * FROM posts")
	defer rows.Close()
	if err != nil {
		return posts, err
	}

	var post Post
	for rows.Next() {
		rows.StructScan(&post)
		posts = append(posts, post)
	}
	return
}

type Post struct {
	ID int64 `db:"id"`
	Text string `db:"text"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}
