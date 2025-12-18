package middleware

import (
	"net/http"

	"github.com/gin-api-demo/utils"
	"github.com/gin-gonic/gin"
)

func JWTUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取Authorization字段，并赋值给post.UserID,格式以Bearer开头
		tokenString := c.GetHeader("Authorization")
		//jwt := fmt.Sprintf("Bearer %s", tokenString)
		// 解析验证token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"msg":    "无效的Token或已被篡改",
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.UserName)
	}
}
