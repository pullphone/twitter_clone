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
	response := entity.ResponseResult{}

	c.Bind(&post)
	id, err := controller.Service.Add(post)
	if err != nil {
		response.Result = false
		response.Data = entity.Error{
			ErrorName: "CANNOT_STORE_POST",
			Message: "cannot store new post",
		}
		c.JSON(500, response)
		return
	}

	post.ID = id
	post.CreatedAt = time.Now().Unix()
	post.UpdatedAt = time.Now().Unix()
	response.Result = true
	response.Data = post
	c.JSON(201, response)
}

func (controller *PostController) Show(c Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	post, err := controller.Service.PostById(id)
	response := entity.ResponseResult{}

	if err != nil {
		response.Result = false
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
		response.Data = error

		c.JSON(code, response)
		return
	}

	response.Result = true
	response.Data = post
	c.JSON(200, response)
}