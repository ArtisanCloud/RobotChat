package main

import (
	"github.com/ArtisanCloud/RobotChat/app/routes/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	api.InitializeAPIRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
