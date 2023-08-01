package api

import (
	chatGLM2 "github.com/ArtisanCloud/RobotChat/app/controller/trainBot/chatGLM"
	"github.com/ArtisanCloud/RobotChat/app/request/chatGLM"
	"github.com/gin-gonic/gin"
)

func InitTrainerBotAPIRoutes(r *gin.Engine) {
	apiTrainerBotRouter := r.Group("/api/v1/bot-trainer")
	{
		apiTrainerBotRouter = apiTrainerBotRouter.Group("/convert")
		// before routes
		{
			apiTrainerBotRouter.POST("/excel-to-self-cognition-data", chatGLM.ValidateConvertExcelToSelfCognitionData, chatGLM2.APIConvertExcelToSelfCognitionData)
		}
	}
}
