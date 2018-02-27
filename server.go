package main

import (
	"github.com/pullphone/twitter_clone/controller"
	"github.com/pullphone/twitter_clone/infrastructure"
)

func main() {
	infrastructure.DBInit()
	defer infrastructure.DBClose()
	controller.Router.Run()
}
