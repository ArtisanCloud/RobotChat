package sd

import (
	request2 "github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/model/controlNet"
	"github.com/gin-gonic/gin"
)

type ParaControlNetDetectInfo struct {
	*controlNet.DetectInfo
}

func ValidateControlNetDetectInfo(c *gin.Context) {
	var params ParaImage2Image

	err := request2.ValidatePara(c, &params)
	if err != nil {
		return
	}

	c.Set("params", &params)
	c.Next()
}
