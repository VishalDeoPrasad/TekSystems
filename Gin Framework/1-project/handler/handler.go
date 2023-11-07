package handler

import (
	"golang/1-project/entity"
	"golang/1-project/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

// FindAll implements VideoController.
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save implements VideoController.
/*func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}*/

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func New(service service.VideoService) VideoController {
	//validate := validator.New()
	//validate.RegisterValidation("is-cool", validators.validateCoolTitle)

	return &controller{
		service: service,
	}
}
