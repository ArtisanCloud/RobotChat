package controller

import (
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APICompletion(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*request.ParaPrompt)

	res, err := service.SrvChatBot.Completion(c.Request.Context(), &param.ChatCompletionRequest)
	if err != nil {
		panic(err)
	}

	response.Success(c, res)
	return

}

func APIChatCompletion(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*request.ParaPrompt)

	err := service.SrvChatBot.ChatCompletion(c.Request.Context(), &param.ChatCompletionRequest)
	if err != nil {
		panic(err)
	}

	response.Success(c, nil)
	return

}
