package chatGLM

import (
	"github.com/ArtisanCloud/RobotChat/app/request"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type ParaConvertExcelToSelfCognitionData struct {
	ExcelFile *multipart.File
}

func ValidateConvertExcelToSelfCognitionData(c *gin.Context) {
	var params ParaConvertExcelToSelfCognitionData

	excelFile, err := request.ValidateFile(c, "excelFile")
	if err != nil {
		return
	}

	params.ExcelFile = excelFile
	c.Set("params", &params)
	c.Next()
}
