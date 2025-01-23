package controllers

import (
	"gin-backend-api/global"
	"gin-backend-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var ExchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&ExchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "json error",
		})
		return
	}
	ExchangeRate.Date = time.Now()

	if err := global.Db.AutoMigrate(&ExchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error inititalzate model",
		})
		return
	}

	if err := global.Db.Create(&ExchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error create exchangeRate",
		})
		return
	}
	ctx.JSON(http.StatusOK, ExchangeRate)
}

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRates []models.ExchangeRate

	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error get exchangeRates,error:" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)
}
