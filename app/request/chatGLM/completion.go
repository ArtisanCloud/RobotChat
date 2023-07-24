package chatGLM

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot/model"
	"github.com/gin-gonic/gin"
)

type ParaCompletion struct {
	model.CompletionRequest
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
