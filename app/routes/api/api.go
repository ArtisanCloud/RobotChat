package api

import (
	"github.com/ArtisanCloud/RobotChat/app/controller"
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/gin-gonic/gin"
)

func InitializeAPIRoutes(r *gin.Engine) {

	InitChatBotAPIRoutes(r)
	InitArtBotAPIRoutes(r)

	apiRobotRouter := r.Group("/api/v1")
	{
		apiRobotRouter.POST("/webhook/queue/notify", request.ValidateQueueNotify, controller.APIQueueNotify)
	}

}
