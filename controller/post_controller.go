package controller

import (
	"github.com/pullphone/twitter_clone/service"
	"github.com/pullphone/twitter_clone/repository"
	"github.com/pullphone/twitter_clone/entity"
	"strconv"
	"time"
	"database/sql"
)

type PostController struct {
	Service service.PostService
}

func NewPostController(sqlHandler repository.SqlHandler) *PostController {
	return &PostController{
		Service: service.PostService{
			PostRepository: &repository.PostRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PostController) Create(c Context) {
	post := entity.Post{}
	c.Bind(&post)
	id, err := controller.Service.Add(post)
	if err != nil {
		c.JSON(500, err)
		return
	}

	post.ID = id
	post.CreatedAt = time.Now().Unix()
	post.UpdatedAt = time.Now().Unix()
	c.JSON(201, post)
}

func (controller *PostController) Show(c Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	post, err := controller.Service.PostById(id)
	if err != nil {
		code := 500
		error := entity.Error{}
		if err == sql.ErrNoRows {
			code = 404
			error.ErrorName = "NOT_FOUND_POSTS"
			error.Message = "not found posts"
		} else {
			error.ErrorName = "CAUGHT_ERROR"
			error.Message = err.Error()
		}
		c.JSON(code, error)
		return
	}
	c.JSON(200, post)
}