package request

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func ValidatePara(context *gin.Context, reqInfo interface{}) (err error) {

	if err = context.ShouldBind(reqInfo); err != nil {
		// 如果发生错误，您可以直接返回响应并中断请求链
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	return err

}

func ValidateFile(context *gin.Context, fileName string) (*multipart.File, error) {

	// 获取上传的文件
	file, _, err := context.Request.FormFile(fileName)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	return &file, nil
}
