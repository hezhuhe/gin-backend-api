package middlewares

import (
	"gin-backend-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			utils.Error(ctx, http.StatusUnauthorized, "error get token")
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)
		if err != nil {
			utils.Error(ctx, http.StatusUnauthorized, "parse token error"+err.Error())
			return
		}
		ctx.Set("username", username)
		ctx.Next()
	}
}
