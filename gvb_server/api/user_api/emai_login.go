package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/ctype"
	"gvb_server/models"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	if FK.Gin.BindJSON(c, &cr) {
		var user models.UserModel
		if !FK_SQL.Is(&user, "user_name = ? or email = ?", cr.UserName, cr.UserName) {
			BY.JSON(c, 1000, "用户名/邮件不存在")
			return
		}
		// 校验密码
		if user.Password != cr.Password {
			BY.JSON(c, 1000, "用户名或密码错误")
			return
		}

		// 邮箱登录
		token, err := JWT.GenToken(JWT.Info{
			UserName: user.UserName,
			NickName: user.NickName,
			Role:     user.Role,
			UserID:   user.ID,
		})

		if BY.JSONErr(c, "token生成失败", err) {
			BY.JSON(c, 0, "登录成功", token)
		}

		ip := FK.Gin.GetAddr(c)
		FK_SQL.Create(&models.LoginDataModel{
			UserID:   user.ID,
			IP:       ip.IP,
			NickName: user.NickName,
			Token:    token,
			Device:   "",
			Addr:     ip.Addr,
			LogType:  ctype.SignEmail,
		})
	}
}
