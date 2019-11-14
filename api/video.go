package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	videoService := service.CreateVideoService{}

	if err := c.ShouldBind(&videoService); err == nil {
		res := videoService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

//ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	videoService := service.ShowVideoService{}

	res := videoService.Show(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

//ListVideo 视频列表接口
func ListVideo(c *gin.Context) {
	videoService := service.ListVideoService{}

	res := videoService.List()
	c.JSON(http.StatusOK, res)
}

//UpdateVideo 修改视频信息
func UpdateVideo(c *gin.Context) {
	videoService := service.UpdateVideoService{}

	if err := c.ShouldBind(&videoService); err == nil {
		res := videoService.Update(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

//DeleteVideo 删除视频信息
func DeleteVideo(c *gin.Context) {
	videoService := service.DeleteVideoService{}

	res := videoService.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
