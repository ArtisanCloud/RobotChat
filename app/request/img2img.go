package request

import (
	"github.com/ArtisanCloud/RobotChat/robots/artBot/request"
	"github.com/gin-gonic/gin"
)

type ParaImage2Image struct {
	ConversationId string `json:"conversationId,optional"`
	SessionId      string `json:"sessionId,optional"`
	JobId          string `json:"jobId,optional"`
	request.Image2Image
}

func ValidateImage2Image(c *gin.Context) {
	var params ParaImage2Image

	err := ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
