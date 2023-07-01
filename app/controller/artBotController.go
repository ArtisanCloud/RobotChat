package controller

import (
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APITxt2Image(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*request.ParaText2Image)

	res, err := service.SrvArtBot.Txt2Image(c.Request.Context(), &param.Text2Image)
	if err != nil {
		panic(err)
	}

	response.Success(c, res)
	return

}

func APIChatTxt2Image(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*request.ParaText2Image)

	ctx := c.Request.Context()

	req := param
	err := service.SrvArtBot.ChatTxt2Image(ctx, req)
	if err != nil {
		panic(err)
	}

	param.Text2Image.Prompt += "-1"
	req = param
	err = service.SrvArtBot.ChatTxt2Image(ctx, req)
	if err != nil {
		panic(err)
	}

	param.Text2Image.Prompt += "-2"
	req = param
	err = service.SrvArtBot.ChatTxt2Image(ctx, req)
	if err != nil {
		panic(err)
	}

	param.Text2Image.Prompt += "-3"
	req = param
	err = service.SrvArtBot.ChatTxt2Image(ctx, req)
	if err != nil {
		panic(err)
	}

	response.Success(c, nil)
	return

}
