package handler

import (
	"golang/1-project/model"
	"golang/1-project/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) model.Video
	FindAll() []model.Video
}

type controller struct {
	service service.VideoService
}

// FindAll implements VideoController.
func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) model.Video {
	var video model.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}
