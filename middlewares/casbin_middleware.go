package middlewares

import (
	"gin-backend-api/utils"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetString("user") // 获取当前用户
		path := c.Request.URL.Path  // 请求路径
		method := c.Request.Method  // 请求方法

		// 获取用户的角色列表
		roles, err := enforcer.GetRolesForUser(user) // 假设你有一个函数获取用户角色
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "获取用户角色失败")
		}
		allowed := false
		for _, role := range roles {
			if ok, _ := enforcer.Enforce(role, path, method); ok {
				allowed = true
				break
			}
		}

		if !allowed {
			utils.Error(c, http.StatusForbidden, "权限不足")
			c.Abort()
			return
		}

		c.Next()
	}
}
