package request

import (
	"github.com/gin-gonic/gin"
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
