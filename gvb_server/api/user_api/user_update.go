package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/INIT/ctype"
	"gvb_server/models"
)

type UserRole struct {
	Role     ctype.Role `json:"role"`
	NickName string     `json:"nick_name"`
	UserID   uint       `json:"user_id" binding:"required" msg:"ID不能为空"`
}

func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	if FK.Gin.BindJSON(c, &cr) {
		var model models.UserModel
		if !FK_SQL.Is(&model, cr.UserID) {
			BY.JSON(c, "ID不存在")
			return
		}
		ok := FK_SQL.Update(&model, map[string]any{
			"role":      cr.Role,
			"nick_name": cr.NickName,
		})
		FK.JSON(c, ok, "更新", cr)
	}
}
