package JWT

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
)

// Auth 登录认证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := FK.Gin.GetHeader(c, "token")
		claims := ParseToken(token)
		if claims == nil {
			BY.JSON(c, []string{token})
			c.Abort()
			return
		}

		c.Set("claims", claims) // 用户信息
	}
}

// Token 返回Token claims
func Token(c *gin.Context) *Claims {
	user, _ := c.Get("claims")
	if user == nil {
		return &Claims{}
	}
	return user.(*Claims)
}
