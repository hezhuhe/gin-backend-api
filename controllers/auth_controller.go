package controllers

import (
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	hashedPwd, err := utils.HashPasswoerd(user.PassWoord)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	user.PassWoord = hashedPwd
}
