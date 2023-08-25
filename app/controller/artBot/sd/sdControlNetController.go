package sd

import (
	"github.com/ArtisanCloud/RobotChat/app/request/sd"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APIGetControlNetModels(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetControlNetModelList(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIGetControlNetModules(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetControlNetModuleList(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIGetControlNetControlTypesList(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetControlNetControlTypesList(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIGetControlNetSettings(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetControlNetSettings(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIGetVersion(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetControlNetVersion(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIDetect(c *gin.Context) {
	ctx := c.Request.Context()
	params, _ := c.Get("params")
	param := params.(*sd.ParaControlNetDetectInfo)

	res, err := service.Michelle.DetectControlNet(ctx, param)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}
