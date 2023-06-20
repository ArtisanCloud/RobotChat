package api

import "github.com/gin-gonic/gin"

func InitChatBotAPIRoutes(r *gin.Engine) {
	apiChatBotRouter := r.Group("/api/v1/chatBot")
	{
		apiChatBotRouter.Use(nil)
		{
			apiChatBotRouter.POST("/createCompletion")

		}
	}
}
