package router

import (
	"backend/controllers"
	"backend/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/article", controllers.CreateArticle)
		api.GET("/article", controllers.GetArticles)
		api.GET("/article/:id", controllers.GetArticleById)

		api.POST("article/:id/like", controllers.LikeArticle)
		api.GET("article/:id/like", controllers.GetArticleLikes)

	}

	return r
}
