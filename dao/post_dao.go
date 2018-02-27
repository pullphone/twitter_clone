package dao

import (
	"github.com/pullphone/twitter_clone/infrastructure"
	"time"
	"database/sql"
)

type PostDao interface {
	Insert(text string) (id int64, err error)
	Get(targetId int64) (postRow Post, err error)
	GetList() (postRows []Post, err error)
}

type postDao struct {
	db infrastructure.DB
}

func NewPostDao() *postDao {
	return &postDao{
		db: infrastructure.GetDatabase("default"),
	}
}

func (dao *postDao) Insert(text string) (id int64, err error) {
	result, err := dao.db.Execute("INSERT INTO posts (text, created_at) VALUES (?, CURRENT_TIMESTAMP)", text)
	if err != nil {
		return
	}

	id, err = result.LastInsertId()
	return
}

func (dao *postDao) Get(targetId int64) (post Post, err error) {
	rows, err := dao.db.Query("SELECT * FROM posts WHERE id = ?", targetId)
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

func (dao *postDao) GetList() (posts []Post, err error) {
	rows, err := dao.db.Query("SELECT * FROM posts")
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
