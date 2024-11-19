package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/ctype"
	"gvb_server/models"
)

// UserListView 用户列表
func (UserApi) UserListView(c *gin.Context) {
	var cr FK_Struct.Mysql
	token := JWT.Token(c)
	if FK.Gin.BindQuery(c, &cr) {
		list := FK_SQL.GetList(models.UserModel{}, cr)
		var List []models.UserModel
		for _, user := range list {
			if token.Role != ctype.PermissionAdmin {
				user.UserName = "权限不足"
				user.Tel = FK.Desensitization.Tel(user.Tel)
				user.Email = FK.Desensitization.Email(user.Email)
			}
			List = append(List, user)
		}
		BY.JSONList(c, len(list), list)
	}
}
