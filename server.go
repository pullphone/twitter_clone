package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pullphone/twitter_clone/controller"
	"github.com/pullphone/twitter_clone/dao"
)

func main() {
	router := gin.Default()

	postController := controller.NewPostController(dao.NewSqlHandler())

	router.POST("/post", func(c *gin.Context) { postController.Create(c) })
	router.GET("/post/:id", func(c *gin.Context) { postController.Show(c) })

	router.Run()
}
