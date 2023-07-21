package api

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/RobotChat/app/controller"
	"github.com/ArtisanCloud/RobotChat/app/middleware"
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeAPIRoutes(r *gin.Engine) {

	r.Use(middleware.PanicMiddleware())

	InitChatBotAPIRoutes(r)
	InitArtBotAPIRoutes(r)
	InitBotTrainerAPIRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, object.HashMap{
			"Greet":       "Hello, This is the Robot Chat butler",
			"Description": "We provide ChatBot Joy and ArtBot Michelle robots, who help you connect many ai channels, like gpt, stable diffusion etc ",
			"Version":     "V1.0.0",
		})
	})

	apiRobotRouter := r.Group("/api/v1")
	{
		apiRobotRouter.POST("/webhook/queue/notify", request.ValidateQueueNotify, controller.APIQueueNotify)
	}

}
