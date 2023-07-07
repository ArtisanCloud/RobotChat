package api

import (
	"github.com/ArtisanCloud/RobotChat/app/controller/artBot/sd"
	"github.com/ArtisanCloud/RobotChat/app/middleware"
	sd2 "github.com/ArtisanCloud/RobotChat/app/request/sd"
	"github.com/gin-gonic/gin"
)

func InitArtBotAPIRoutes(r *gin.Engine) {
	apiArtBotRouter := r.Group("/api/v1/artBot")
	{
		apiArtBotRouter.Use(middleware.ArtBotIsAwaken)
		{
			apiSDRouter := apiArtBotRouter.Group("/sd")
			// before routes
			{
				apiSDRouter.POST("/txt2img", sd2.ValidateText2Image, sd.APITxt2Image)
				apiSDRouter.POST("/img2img", sd2.ValidateImage2Image, sd.APIImage2Image)
				apiSDRouter.POST("/chat/txt2img", sd2.ValidateText2Image, sd.APIChatTxt2Image)
				apiSDRouter.POST("/chat/img2img", sd2.ValidateImage2Image, sd.APIChatImage2Image)

				// progress
				apiSDRouter.POST("/progress", sd.APIProgress)

				// option
				apiSDRouter.GET("/options", sd.APIGetOptions)
				apiSDRouter.POST("/options", sd2.ValidateSetOptions, sd.APIPostOptions)
			}
		}

	}
}
