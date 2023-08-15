package sd

import (
	"github.com/ArtisanCloud/RobotChat/app/request/sd"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APIGetModels(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetModels(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIGetSamplers(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetSamplers(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIProgress(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.Progress(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIGetOptions(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := service.Michelle.GetOptions(ctx)

	if err != nil {
		panic(err)
	}
	response.Success(c, res)
	return
}

func APIPostOptions(c *gin.Context) {
	params, _ := c.Get("params")
	param := params.(*sd.ParaSetOptions)

	ctx := c.Request.Context()

	err := service.Michelle.SetOptions(ctx, param.OptionsRequest)

	if err != nil {
		panic(err)
	}
	response.Success(c, nil)
	return
}
