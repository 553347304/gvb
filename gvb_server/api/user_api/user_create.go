package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/INIT/ctype"
	"gvb_server/models"
)

type UserModel struct {
	NickName string     `json:"nick_name" binding:"required" msg:"不能为空"`  // 昵称
	UserName string     ` json:"user_name" binding:"required" msg:"不能为空"` // 用户名
	Password string     `json:"password" binding:"required" msg:"不能为空"`   // 密码 	// 头像id
	PassWord string     `json:"pass_word" binding:"required" msg:"不能为空"`
	Email    string     `json:"email" binding:"required" msg:"不能为空"` // 邮箱
	Tel      string     ` json:"tel" binding:"required" msg:"不能为空"`  // 手机号
	Role     ctype.Role ` json:"role" binding:"required" msg:"不能为空"` // 权限  1 管理员  2 普通用户  3 游客
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserModel
	if FK.Gin.BindJSON(c, &cr) {
		if cr.Password != cr.PassWord {
			BY.JSON(c, "注册失败: 两次密码不一致")
			return
		}
		var model models.UserModel
		if FK_SQL.Is(&model, "user_name = ?", cr.UserName) {
			BY.JSON(c, "注册失败: 用户名已存在")
			return
		}
		ip := FK.Gin.GetAddr(c)
		ok := FK_SQL.Create(&models.UserModel{
			NickName:   cr.NickName,
			UserName:   cr.UserName,
			Password:   cr.Password,
			Email:      cr.Email,
			Tel:        cr.Tel,
			Token:      "token",
			Role:       cr.Role,
			Avatar:     Avatar,
			IP:         ip.IP,
			Addr:       ip.Addr,
			SignStatus: 4,
		})
		FK.JSON(c, ok, "注册用户", ok)
	}
}
