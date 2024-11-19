package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/models"
)

type UserUpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     // 新密码
}

func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UserUpdatePasswordRequest
	if FK.Gin.BindJSON(c, &cr) {
		_user, _ := c.Get("user")
		user := _user.(*JWT.Claims)
		var model models.UserModel
		if !FK_SQL.Is(&model, user.UserID) {
			BY.JSON(c, "用户不存在")
			return
		}
		if model.Password != cr.OldPwd {
			BY.JSON(c, "密码错误")
			return
		}
		ok := FK_SQL.Update(&model, map[string]any{
			"password": cr.Pwd,
		})
		FK.JSON(c, ok, "修改密码", ok)
	}

}
