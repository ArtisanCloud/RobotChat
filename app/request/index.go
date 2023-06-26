package request

import "github.com/gin-gonic/gin"

func ValidatePara(context *gin.Context, reqInfo interface{}) (err error) {

	if err = context.ShouldBind(reqInfo); err != nil {
		if err = context.ShouldBindJSON(reqInfo); err != nil {

		}
	}
	return err

}
