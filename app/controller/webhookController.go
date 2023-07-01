package controller

import (
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APIQueueNotify(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*request.ParaQueueNotify)

	service.SrvArtBot.WebhookText(c.Request.Context(), param)

	response.Success(c, param)
	return

}
