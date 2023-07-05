package openai

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/request"
	"github.com/gin-gonic/gin"
)

type ParaCompletion struct {
	ConversationId string `json:"conversationId, optional"`
	SessionId      string `json:"sessionId, optional"`
	JobId          string `json:"jobId, optional"`
	request.CompletionRequest
}

func ValidateCompletion(c *gin.Context) {
	var params ParaCompletion

	err := request2.ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
