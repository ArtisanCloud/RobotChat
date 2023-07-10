package sd

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/gin-gonic/gin"
)

type ParaImage2Image struct {
	ConversationId string `json:"conversationId,optional"`
	SessionId      string `json:"sessionId,optional"`
	JobId          string `json:"jobId,optional"`
	*model.Image2ImageRequest
}

func ValidateImage2Image(c *gin.Context) {
	var params ParaImage2Image

	err := request2.ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
