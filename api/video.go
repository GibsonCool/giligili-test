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
