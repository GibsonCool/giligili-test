package service

import (
	"giligili/model"
	"giligili/serializer"
)

//UpdateVideoService 更新视频信息的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=8,max=40"`
}

//Update 更新视频信息
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	// 先查询
	e := model.DB.First(&video, id).Error
	if e != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: e.Error(),
		}
	}

	// 修改
	video.Title = service.Title
	video.Info = service.Info
	e = model.DB.Save(&video).Error
	if e != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "视频更新失败",
			Error: e.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
