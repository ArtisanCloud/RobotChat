package sd

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model"
	"github.com/gin-gonic/gin"
)

type ParaSetOptions struct {
	*model.OptionsRequest
}

func ValidateSetOptions(c *gin.Context) {
	var params ParaSetOptions

	err := request2.ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
