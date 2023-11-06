package service

import "golang/1-project/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func (vs *videoService) Save(video entity.Video) entity.Video {
	vs.videos = append(vs.videos, video)
	return video
}

func (vs *videoService) FindAll() []entity.Video {
	return vs.videos
}

//The New function creates and returns an instance of videoService as a VideoService.
func New() VideoService {
	return &videoService{}
}
