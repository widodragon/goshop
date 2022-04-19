package service

import "github.com/widodragon/goshop/entity"

type VideoService interface {
	FindAll() []entity.Video
	Save(video entity.Video) entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) FindAll() []entity.Video {
	var emptyArr = []entity.Video{}
	if service.videos == nil {
		return emptyArr
	}
	return service.videos
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}
