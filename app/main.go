package main

import (
	"github.com/ArtisanCloud/RobotChat/app/routes/api"
	"github.com/ArtisanCloud/RobotChat/app/service"
	fmt "github.com/ArtisanCloud/RobotChat/pkg/printx"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	config := rcconfig.LoadRCConfig()

	// 初始化服务层
	service.InitService(config)

	r := gin.Default()
	api.InitializeAPIRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Dump(err)
	}
}
