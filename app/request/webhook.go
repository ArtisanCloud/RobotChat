package request

import (
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/gin-gonic/gin"
)

type ParaQueueNotify struct {
	*model.Job
}

func ValidateQueueNotify(c *gin.Context) {
	var params ParaQueueNotify

	err := ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
