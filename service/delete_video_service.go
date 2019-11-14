package service

import (
	"giligili/model"
	"giligili/serializer"
)

//DeleteVideoService 视频详情的服务
type DeleteVideoService struct {
}

//Create 展示视频
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	// 先查询
	e := model.DB.First(&video, id).Error
	if e != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "要删除的视频不存在",
			Error: e.Error(),
		}
	}

	// 后删除
	e = model.DB.Delete(&video).Error
	if e != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频删除失败",
			Error: e.Error(),
		}
	}

	return serializer.Response{
		Msg: "删除成功",
	}
}
