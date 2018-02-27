package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pullphone/twitter_clone/service"
)

var Router *gin.Engine

var postController = &PostController{
	postService: service.NewPostService(),
}

func init() {
	Router = gin.Default()
	Router.Use(maxAllowed(25))

	Router.POST("/post", func(c *gin.Context) { postController.Post(c) })
	Router.GET("/post/:id", func(c *gin.Context) { postController.Get(c) })
	Router.GET("/posts", func(c *gin.Context) { postController.GetAll(c) })

	Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"result": false,
			"data": gin.H{
				"error_name": "ROUTE_NOT_FOUND",
				"message":    "route not found",
			},
		})
	})
}

func maxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire() // before request
		defer release() // after request
		c.Next()

	}
}
