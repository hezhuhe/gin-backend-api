package main

import (
	"backend/config"
	"backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	// fmt.Println(config.AppConfig.App.Name)
	// fmt.Println(config.AppConfig.App.Port)
	type SuccessRespons struct {
		Code int
		Msg  string
	}
	infoTest := SuccessRespons{
		Code: 200,
		Msg:  "success",
	}
	r := router.SetupRouter()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, infoTest)
	})

	// addr := "10.22.40.90" + config.AppConfig.App.Port
	r.Run(config.AppConfig.App.Port) // 监听并在 0.0.0.0:8080 上启动服务
}
