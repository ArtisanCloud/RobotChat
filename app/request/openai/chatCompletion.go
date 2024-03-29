package openai

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/model"
	"github.com/gin-gonic/gin"
)

type ParaChatCompletion struct {
	ConversationId string `json:"conversationId,optional"`
	SessionId      string `json:"sessionId,optional"`
	JobId          string `json:"jobId,optional"`
	model.ChatCompletionRequest
}

func ValidateChatCompletion(c *gin.Context) {
	var params ParaChatCompletion

	err := request2.ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
