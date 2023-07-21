package api

import (
	chatGLM2 "github.com/ArtisanCloud/RobotChat/app/controller/botTrainer/chatGLM"
	"github.com/ArtisanCloud/RobotChat/app/request/chatGLM"
	"github.com/gin-gonic/gin"
)

func InitBotTrainerAPIRoutes(r *gin.Engine) {
	apiBotTrainerRouter := r.Group("/api/v1/bot-trainer")
	{
		apiBotTrainerRouter = apiBotTrainerRouter.Group("/convert")
		// before routes
		{
			apiBotTrainerRouter.POST("/excel-to-self-cognition-data", chatGLM.ValidateConvertExcelToSelfCognitionData, chatGLM2.APIConvertExcelToSelfCognitionData)
		}
	}
}
