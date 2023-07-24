package api

import (
	"github.com/ArtisanCloud/RobotChat/app/controller/chatBot/chatGLM"
	"github.com/ArtisanCloud/RobotChat/app/controller/chatBot/openai"
	"github.com/ArtisanCloud/RobotChat/app/middleware"
	openai2 "github.com/ArtisanCloud/RobotChat/app/request/openai"
	"github.com/gin-gonic/gin"
)

func InitChatBotAPIRoutes(r *gin.Engine) {
	apiChatBotRouter := r.Group("/api/v1/chat-bot")
	{
		apiChatBotRouter.Use(middleware.ChatBotIsAwaken)
		{
			apiOpenAIRouter := apiChatBotRouter.Group("/openai")
			{
				apiOpenAIRouter.POST("/completion", openai2.ValidateCompletion, openai.APICompletion)
				apiOpenAIRouter.POST("/chat/completion", openai2.ValidateChatCompletion, openai.APIChatCompletion)
				apiOpenAIRouter.POST("/stream/completion", openai2.ValidateCompletion, openai.APIStreamCompletion)
			}

			apiChatGLMRouter := apiChatBotRouter.Group("/glm")
			{
				apiChatGLMRouter.POST("/completion", openai2.ValidateCompletion, chatGLM.APICompletion)
				apiChatGLMRouter.POST("/chat/completion", openai2.ValidateChatCompletion, chatGLM.APIChatCompletion)
				apiChatGLMRouter.POST("/stream/completion", openai2.ValidateCompletion, chatGLM.APIStreamCompletion)
			}
		}
	}
}
