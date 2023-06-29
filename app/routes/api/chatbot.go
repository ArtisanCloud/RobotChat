package api

import (
	"github.com/ArtisanCloud/RobotChat/app/controller"
	"github.com/ArtisanCloud/RobotChat/app/middleware"
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/gin-gonic/gin"
)

func InitChatBotAPIRoutes(r *gin.Engine) {
	apiChatBotRouter := r.Group("/api/v1/chatBot")
	{
		apiChatBotRouter.Use(middleware.ChatBotIsAwaken)
		{
			apiChatBotRouter.POST("/completion", request.ValidatePrompt, controller.APICompletion)
			apiChatBotRouter.POST("/chat/completion", request.ValidatePrompt, controller.APIChatCompletion)
		}
	}
}
