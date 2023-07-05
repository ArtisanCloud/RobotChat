package sd

import (
	"github.com/ArtisanCloud/RobotChat/app/request/sd"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APITxt2Image(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*sd.ParaText2Image)

	res, err := service.Michelle.Txt2Image(c.Request.Context(), &param.Text2Image)
	if err != nil {
		panic(err)
	}

	response.Success(c, res)
	return

}

func APIImage2Image(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*sd.ParaImage2Image)

	res, err := service.Michelle.Image2Image(c.Request.Context(), &param.Image2Image)
	if err != nil {
		panic(err)
	}

	response.Success(c, res)
	return

}

func APIChatTxt2Image(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*sd.ParaText2Image)

	ctx := c.Request.Context()

	req := param
	job, err := service.Michelle.ChatTxt2Image(ctx, req)
	if err != nil {
		panic(err)
	}
	response.Success(c, job)
	return

}

func APIChatImage2Image(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*sd.ParaImage2Image)

	ctx := c.Request.Context()

	req := param
	job, err := service.Michelle.ChatImage2Image(ctx, req)
	if err != nil {
		panic(err)
	}
	response.Success(c, job)
	return

}
