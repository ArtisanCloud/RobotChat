package api

import (
	"github.com/ArtisanCloud/RobotChat/app/controller"
	"github.com/ArtisanCloud/RobotChat/app/middleware"
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/gin-gonic/gin"
)

func InitArtBotAPIRoutes(r *gin.Engine) {
	apiArtBotRouter := r.Group("/api/v1/artBot")
	{
		// before routes
		apiArtBotRouter.Use(middleware.ArtBotIsAwaken)
		{
			apiArtBotRouter.POST("/txt2img", request.ValidateTxt2Image, controller.APITxt2Image)
			apiArtBotRouter.POST("/chat/txt2img", request.ValidateTxt2Image, controller.APIChatTxt2Image)

		}

	}
}
