package chatGLM

import (
	"github.com/ArtisanCloud/RobotChat/app/request/chatGLM"
	"github.com/ArtisanCloud/RobotChat/app/response"
	"github.com/ArtisanCloud/RobotChat/app/service"
	"github.com/ArtisanCloud/RobotChat/pkg/objectx"
	"github.com/gin-gonic/gin"
)

func APIConvertExcelToSelfCognitionData(c *gin.Context) {

	params, _ := c.Get("params")
	param := params.(*chatGLM.ParaConvertExcelToSelfCognitionData)

	savePath, err := objectx.GetSavePath("temp/trainer/data_format/", "convert", "json")
	if err != nil {
		panic(err)
	}

	err = service.Michael.ConvertExcelToSelfCognitionData(c.Request.Context(), param.ExcelFile, savePath)
	if err != nil {
		panic(err)
	}

	response.Success(c, err)
	return

}
