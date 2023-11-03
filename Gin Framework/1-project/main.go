package main

import (
	"golang/1-project/handler"
	"golang/1-project/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService    = service.New()
	videoController handler.VideoController = handler.New(videoService)
)

func main() {
	r := gin.Default()

	r.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	r.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	r.Run(":8080")

}

/*
	The code you've provided is for a simple web application written in Go that uses
	the Gin framework. This application allows you to interact with video data through
	HTTP requests. It's structured with a basic Model-Service-Controller architecture.
*/
