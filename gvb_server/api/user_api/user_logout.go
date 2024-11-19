package user_api

import (
	"github.com/gin-gonic/gin"
	BY "gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/global"
	"time"
)

func (by UserApi) UserLogoutView(c *gin.Context) {
	// 过期时间
	token := c.Request.Header.Get("token")
	claims := JWT.Token(c)
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	head := "logout_"
	err := global.Redis.Set(head+token, "", diff).Err()
	if err != nil {
		BY.JSON(c, "注销失败")
		return
	}
	// 是否存在
	keys := global.Redis.Keys(head + "*").Val()
	if !BY.Map.IsList(head+token, keys) {
		BY.JSON(c, "token已注销")
		c.Abort()
		return
	}
}
