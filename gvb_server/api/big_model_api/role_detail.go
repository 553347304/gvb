package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

type TagResponse struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type RoleDetailResponse struct {
	models.MODEL
	Id        uint          `json:"id"`
	Name      string        `json:"name"`
	Icon      string        `json:"icon"`
	Abstract  string        `json:"abstract"`
	Tags      []TagResponse `json:"tags"`
	ChatTotal int           `json:"chat_total"`
}

func (BigModelApi) RoleDetailView(c *gin.Context) {
	var cr FK_Struct.UID
	if FK.Gin.BindURI(c, &cr) {
		var role models.BigModelRoleModel
		if !FK_SQL.Preload("Tags", &role, cr.ID) {
			BY.JSON(c, 1000, "角色不存在")
			return
		}
		var tagList = make([]TagResponse, 0)
		for _, tag := range role.Tags {
			tagList = append(tagList, TagResponse{
				Id:    tag.ID,
				Title: tag.Title,
				Color: tag.Color,
			})
		}
		var response = RoleDetailResponse{
			MODEL:    role.MODEL,
			Name:     role.Name,
			Icon:     role.Icon,
			Abstract: role.Abstract,
			Tags:     tagList,
		}
		list := FK_SQL.Where(&models.BigModelChatModel{}, "role_id = ?", cr.ID)
		response.ChatTotal = len(list)
		BY.JSON(c, response)
	}
}
