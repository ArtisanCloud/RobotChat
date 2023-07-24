package chatGLM

import (
	"github.com/ArtisanCloud/RobotChat/app/request/openai"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/gin-gonic/gin"
)

func APICompletion(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*openai.ParaCompletion)

	res, err := service.Joy.Completion(c.Request.Context(), &param.CompletionRequest)
	if err != nil {
		panic(err)
	}

	response.Success(c, res)
	return

}

func APIChatCompletion(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*openai.ParaChatCompletion)

	res, err := service.Joy.ChatCompletion(c.Request.Context(), &param.ChatCompletionRequest)
	if err != nil {
		panic(err)
	}

	response.Success(c, res)
	return

}

func APIStreamCompletion(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*openai.ParaCompletion)

	res, err := service.Joy.SteamCompletion(c.Request.Context(), &param.CompletionRequest)
	if err != nil {
		panic(err)
	}

	response.Success(c, res)
	return

}
