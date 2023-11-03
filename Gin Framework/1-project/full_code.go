package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Video represents a video object.
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

var videos []Video

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define a route to get all videos
	r.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videos)
	})

	// Define a route to save a new video
	r.POST("/videos", func(c *gin.Context) {
		var newVideo Video
		c.BindJSON(&newVideo) // Read video data from the request body

		videos = append(videos, newVideo) // Add the new video to the list

		c.JSON(http.StatusCreated, newVideo) // Respond with the saved video
	})

	// Start the server
	r.Run(":8080")
}
