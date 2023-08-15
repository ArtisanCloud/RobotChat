package sd

import (
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APIGetLoras(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetLoras(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIRefreshLoras(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetModels(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}
