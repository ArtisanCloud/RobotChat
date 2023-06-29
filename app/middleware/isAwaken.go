package middleware

import (
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func ArtBotIsAwaken(c *gin.Context) {
	err := service.SrvArtBot.IsAwaken(c.Request.Context())
	if err != nil {
		panic(err)
	}

	c.Next()
}

func ChatBotIsAwaken(c *gin.Context) {

	err := service.SrvChatBot.IsAwaken(c.Request.Context())
	if err != nil {
		panic(err)
	}

	c.Next()
}
