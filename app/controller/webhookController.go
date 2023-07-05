package controller

import (
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/gin-gonic/gin"
)

func APIQueueNotify(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*request.ParaQueueNotify)

	fmt.Dump(param)
	service.Michelle.WebhookText(c.Request.Context(), param)

	response.Success(c, param)
	return

}
