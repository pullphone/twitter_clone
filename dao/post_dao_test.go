package dao

import (
	"database/sql"
	"testing"
)

var dao = NewPostDao()

func TestGet(t *testing.T) {
	post, err := dao.Get(1)
	if err != nil {
		t.Fatal("cannot get")
	}

	if post.ID != 1 {
		t.Fatal("not match post")
	}

	_, err = dao.Get(99999)
	if err != sql.ErrNoRows {
		t.Fatal("not match error")
	}
}

func TestGetList(t *testing.T) {
	posts, err := dao.GetList()
	if err != nil {
		t.Fatal("cannot get_list")
	}

	if len(posts) < 1 {
		t.Fatal("cannot get_list")
	}
}

func TestInsert(t *testing.T) {
	text := "text_test"
	id, err := dao.Insert(text)
	if err != nil {
		t.Fatal("cannot insert")
	}

	if id < 1 {
		t.Fatal("cannot insert")
	}

	post, err := dao.Get(id)
	if post.ID != id {
		t.Fatal("not match inserted post")
	}
	if post.Text != text {
		t.Fatal("not match inserted post")
	}
}