package api

import (
	"github.com/ArtisanCloud/RobotChat/app/controller/artBot/sd"
	"github.com/ArtisanCloud/RobotChat/app/middleware"
	sd2 "github.com/ArtisanCloud/RobotChat/app/request/sd"
	"github.com/gin-gonic/gin"
)

func InitArtBotAPIRoutes(r *gin.Engine) {
	apiArtBotRouter := r.Group("/api/v1/art-bot")
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

				// art models
				apiSDRouter.GET("/model/list", sd.APIGetModels)

				// art samplers
				apiSDRouter.GET("/sampler/list", sd.APIGetSamplers)

				// art Lora
				apiSDRouter.GET("/lora/list", sd.APIGetLoras)
				apiSDRouter.POST("/lora/refresh", sd.APIRefreshLoras)

				// art controlnet
				apiSDRouter.GET("/controlnet/model/list", sd.APIGetControlNetModels)
				apiSDRouter.GET("/controlnet/module/list", sd.APIGetControlNetModules)
				apiSDRouter.GET("/controlnet/control-type/list", sd.APIGetControlNetControlTypesList)
				apiSDRouter.GET("/controlnet/settings", sd.APIGetControlNetSettings)
				apiSDRouter.GET("/controlnet/version", sd.APIGetVersion)
				apiSDRouter.POST("/controlnet/detect", sd.APIDetect)

				// progress
				apiSDRouter.GET("/progress", sd.APIProgress)

				// option
				apiSDRouter.GET("/options", sd.APIGetOptions)
				apiSDRouter.POST("/options", sd2.ValidateSetOptions, sd.APIPostOptions)
			}
		}

	}
}
