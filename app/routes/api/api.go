package api

import "github.com/gin-gonic/gin"

func InitializeAPIRoutes(r *gin.Engine) {

	InitChatBotAPIRoutes(r)
	InitArtBotAPIRoutes(r)

}
