package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否存在错误
		if err, exists := c.Get("error"); exists {
			// 处理错误并返回响应
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			c.Abort() // 终止请求链
		}
	}
}
