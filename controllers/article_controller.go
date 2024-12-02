package controllers

import (
	"backend/global"
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func CreateArticle(ctx *gin.Context) {
	var Article models.Article
	if err := ctx.ShouldBindJSON(&Article); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "json error",
		})
		return
	}

	if err := global.Db.AutoMigrate(&Article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error inititalzate model",
		})
		return
	}
	if err := global.Db.Create(&Article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error create article",
		})
		return
	}
	if err := global.RedisDB.Del(ctx.Request.Context(), cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Article)
}

var cacheKey = "articles"

func GetArticles(ctx *gin.Context) {

	cachedData, err := global.RedisDB.Get(ctx.Request.Context(), cacheKey).Result()

	if err == redis.Nil {
		var articles []models.Article

		if err := global.Db.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		articleJSON, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := global.RedisDB.Set(ctx.Request.Context(), cacheKey, articleJSON, 10*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, articles)

	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		var articles []models.Article

		if err := json.Unmarshal([]byte(cachedData), &articles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, articles)
	}
}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	var article models.Article
	if err := global.Db.Where("id=?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "error get article form database",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, article)
}
