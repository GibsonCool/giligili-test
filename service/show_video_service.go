package service

import (
	"giligili/model"
	"giligili/serializer"
)

//ShowVideoService 视频详情的服务
type ShowVideoService struct {
}

//Create 展示视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	e := model.DB.First(&video, id).Error
	if e != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: e.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
