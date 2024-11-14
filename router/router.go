package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// auth := r.Group("/api/auth")
	// {
	// 	auth.POST("/login", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"msg": "login success",
	// 		})
	// 	})
	// 	auth.POST("/register", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"msg": "register success", // 修正了拼写错误
	// 		})
	// 	})
	// }
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)

		auth.POST("/register", controllers.Register)
	}
	return r
}
