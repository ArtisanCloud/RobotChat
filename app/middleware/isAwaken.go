package middleware

import (
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func ArtBotIsAwaken(c *gin.Context) {
	err := service.Michelle.IsAwaken(c.Request.Context())
	if err != nil {
		panic(err)
	}

	c.Next()
}

func ChatBotIsAwaken(c *gin.Context) {

	err := service.Joy.IsAwaken(c.Request.Context())
	if err != nil {
		panic(err)
	}

	c.Next()
}
