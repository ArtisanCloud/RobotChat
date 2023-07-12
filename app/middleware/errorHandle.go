package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PanicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 处理 panic，可以记录日志或返回自定义错误信息
				log.Println("Panic occurred:", r)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()

		c.Next()
	}
}

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
