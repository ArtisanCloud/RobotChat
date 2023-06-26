package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	"net/http"
	"runtime/debug"
)

// ErrorMiddleware 是一个封装错误处理的中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 处理错误，例如记录日志、返回错误响应等
				_, _ = pretty.Printf("%v\n%s", err, debug.Stack())
				// 返回json结果
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
