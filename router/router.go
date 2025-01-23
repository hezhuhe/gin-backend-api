package router

import (
	"gin-backend-api/controllers"
	"time"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)

		auth.POST("/register", controllers.Register)
	}

	// api := r.Group("/api")
	// api.GET("/exchangeRates", controllers.GetExchangeRate)
	// api.Use(middlewares.AuthMiddleware())
	// {
	// 	api.POST("/exchangeRates", controllers.CreateExchangeRate)
	// 	api.POST("/article", controllers.CreateArticle)
	// 	api.GET("/article", controllers.GetArticles)
	// 	api.GET("/article/:id", controllers.GetArticleById)

	// 	api.POST("article/:id/like", controllers.LikeArticle)
	// 	api.GET("article/:id/like", controllers.GetArticleLikes)

	// }
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
