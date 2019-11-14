package service

import (
	"giligili/model"
	"giligili/serializer"
)

//ListVideoService 视频列表的服务
type ListVideoService struct {
}

//List 展示视频
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	e := model.DB.Find(&videos).Error
	if e != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频列表查询失败",
			Error: e.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
