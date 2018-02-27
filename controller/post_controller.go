package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pullphone/twitter_clone/service"
	"github.com/pullphone/twitter_clone/entity"
	"strconv"
	"time"
	"database/sql"
)

type PostController struct {
	Service service.PostService
}

func (controller *PostController) Post(c *gin.Context) {
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

	post.ID = strconv.FormatInt(id, 10)
	post.CreatedAt = time.Now().Unix()
	post.UpdatedAt = time.Now().Unix()
	response.Result = true
	response.Data = post
	c.JSON(201, response)
}

func (controller *PostController) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	post, err := controller.Service.PostById(id)
	response := entity.ResponseResult{}

	if err != nil {
		response.Result = false
		code := 500
		errorEnt := entity.Error{}
		if err == sql.ErrNoRows {
			code = 404
			errorEnt.ErrorName = "NOT_FOUND_POSTS"
			errorEnt.Message = "not found posts"
		} else {
			errorEnt.ErrorName = "CAUGHT_ERROR"
			errorEnt.Message = err.Error()
		}
		response.Data = errorEnt

		c.JSON(code, response)
		return
	}

	response.Result = true
	response.Data = post
	c.JSON(200, response)
}

func (controllter *PostController) GetAll(c *gin.Context) {
	posts, err := controllter.Service.Posts()
	response := entity.ResponseResult{}

	if err != nil || len(posts) < 1 {
		response.Result = false
		response.Data = entity.Error{
			ErrorName: "NOT_FOUND_POSTS",
			Message: "not posted yet",
		}
		c.JSON(404, response)
		return
	}

	response.Result = true
	response.Data = gin.H{"list": posts}
	c.JSON(200, response)
}
