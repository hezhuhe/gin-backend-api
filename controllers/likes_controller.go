package controllers

import (
	"gin-backend-api/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")
	likeKey := "article" + articleId + ":likes"

	if err := global.RedisDB.Incr(ctx.Request.Context(), likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error insert redis",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": "Successfully liked the article",
	})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := "article:" + articleID + ":likes"

	likes, err := global.RedisDB.Get(ctx.Request.Context(), likeKey).Result()

	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
