package service

import "golang/1-project/model"

type VideoService interface {
	Save(model.Video) model.Video
	FindAll() []model.Video
}

type videoService struct {
	videos []model.Video
}

func (vs *videoService) Save(video model.Video) model.Video {
	vs.videos = append(vs.videos, video)
	return video
}

func (vs *videoService) FindAll() []model.Video {
	return vs.videos
}

//The New function creates and returns an instance of videoService as a VideoService.
func New() VideoService {
	return &videoService{}
}
