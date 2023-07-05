package sd

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/request"
	"github.com/gin-gonic/gin"
)

type ParaText2Image struct {
	ConversationId string `json:"conversationId,optional"`
	SessionId      string `json:"sessionId,optional"`
	JobId          string `json:"jobId,optional"`
	request.Text2Image
}

func ValidateText2Image(c *gin.Context) {
	var params ParaText2Image

	err := request2.ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
