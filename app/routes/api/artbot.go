package api

import "github.com/gin-gonic/gin"

func InitArtBotAPIRoutes(r *gin.Engine) {
	apiArtBotRouter := r.Group("/api/v1/ArtBot")
	{
		apiArtBotRouter.Use(nil)
		{
			apiArtBotRouter.POST("/txt2img")

		}
	}
}
