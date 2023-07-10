package sd

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/gin-gonic/gin"
)

type ParaText2Image struct {
	ConversationId string `json:"conversationId,optional"`
	SessionId      string `json:"sessionId,optional"`
	JobId          string `json:"jobId,optional"`
	*model.Text2ImageRequest
}

func ValidateText2Image(c *gin.Context) {
	var params ParaText2Image

	_ = request2.ValidatePara(c, &params)

	c.Set("params", &params)
	c.Next()
}
