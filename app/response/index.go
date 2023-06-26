package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(context *gin.Context, data interface{}) {

	context.JSON(http.StatusOK, data)
}
