package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

func (BigModelApi) RoleUserHistoryListView(c *gin.Context) {
	claims := JWT.Token(c)

	var roleIdList []uint
	FK_SQL.Group(models.BigModelSessionModel{}, FK_Struct.Mysql{
		Where: []any{"user_id = ?", claims.UserID},
	}, &roleIdList, "role_id")

	var roleList []models.BigModelRoleModel
	global.DB.Order("created_at desc").Find(&roleList, "id in ?", roleIdList)
	var list = make([]RoleItem, 0)
	for _, m := range roleList {
		list = append(list, RoleItem{
			Id:       m.ID,
			Name:     m.Name,
			Abstract: m.Abstract,
			Icon:     m.Icon,
		})
	}
	BY.JSONList(c, len(list), list)
}
