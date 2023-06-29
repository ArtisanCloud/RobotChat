package request

import (
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/request"
	"github.com/gin-gonic/gin"
)

type ParaPrompt struct {
	ConversationId string `json:"conversationId, optional"`
	SessionId      string `json:"sessionId, optional"`
	JobId          string `json:"jobId, optional"`
	request.ChatCompletionRequest
}

func ValidatePrompt(c *gin.Context) {
	var params ParaPrompt

	err := ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
