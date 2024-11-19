package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

type RoleItem struct {
	Id       uint   `json:"id"` // 角色ID
	Name     string `json:"name"`
	Abstract string `json:"abstract"`
	Icon     string `json:"icon"`
}

type TagRoleListResponse struct {
	Id       uint       `json:"id"` // 标签ID
	Title    string     `json:"title"`
	RoleList []RoleItem `json:"role_list"` // 角色列表
}

// TagRoleListView 角色广场
func (BigModelApi) TagRoleListView(c *gin.Context) {
	var _list []models.BigModelTagModel
	global.DB.Preload("Roles").Find(&_list)
	var list = make([]TagRoleListResponse, 0)
	for _, m := range _list {
		roleList := make([]RoleItem, 0)
		for _, role := range m.Roles {
			roleList = append(roleList, RoleItem{
				Id:       role.ID,
				Name:     role.Name,
				Abstract: role.Abstract,
				Icon:     role.Icon,
			})
		}
		list = append(list, TagRoleListResponse{
			Id:       m.ID,
			Title:    m.Title,
			RoleList: roleList,
		})
	}
	BY.JSONList(c, len(list), list)
}
