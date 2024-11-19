package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

type TagUpdateRequest struct {
	Id    uint   `json:"id"`
	Title string `json:"title" binding:"required" msg:"title不能为空"`
	Color string `json:"color" binding:"required" msg:"color不能为空"`
}

func (BigModelApi) TagCreateView(c *gin.Context) {
	var cr TagUpdateRequest
	if FK.Gin.BindJSON(c, &cr) {
		if FK_SQL.Is(&models.BigModelTagModel{}, "title = ?", cr.Title) {
			BY.JSON(c, 1000, "名字重复")
			return
		}
		create := models.BigModelTagModel{
			Title: cr.Title,
			Color: cr.Color,
		}
		R := FK_SQL.Create(&create)
		FK.JSON(c, R, "添加", create.ID)
	}
}

func (BigModelApi) TagUpdateView(c *gin.Context) {
	var cr TagUpdateRequest
	if FK.Gin.BindJSON(c, &cr) {
		var model models.BigModelTagModel
		if !FK_SQL.Is(&model, cr.Id) {
			BY.JSON(c, 1000, "记录不存在")
			return
		}
		R := FK_SQL.Update(&model, BY.Map.StructMapDelEmpty(cr))
		FK.JSON(c, R, "更新数据", R)
	}
}

type TagListResponse struct {
	models.MODEL
	Title     string `json:"title"`
	Color     string `json:"color"`
	RoleTotal int    `json:"role_total"` // 角色个数
}

func (BigModelApi) TagListView(c *gin.Context) {
	var o FK_Struct.Mysql
	if FK.Gin.BindQuery(c, &o) {
		result := FK_SQL.GetList(&models.BigModelTagModel{}, o)
		var list = make([]TagListResponse, 0)
		for _, m := range result {
			list = append(list, TagListResponse{
				MODEL:     m.MODEL,
				Title:     m.Title,
				Color:     m.Color,
				RoleTotal: len(m.Roles),
			})
		}
		BY.JSONList(c, len(list), list)
	}
}

func (BigModelApi) TagRemoveView(c *gin.Context) {
	var id FK_Struct.UID
	if FK.Gin.BindJSON(c, &id) {
		var list []models.BigModelTagModel
		if !FK_SQL.Preload("Roles", &list, id.IDList) {
			BY.JSON(c, 1000, "记录不存在")
			return
		}
		for _, i2 := range list {
			FK_SQL.Association(&i2, "Roles").Delete(i2.Roles) // 删除关联关系
		}
		ok := FK_SQL.Delete(&list)
		FK.JSON(c, ok, "删除标签")
	}
}

func (BigModelApi) TagOptionsView(c *gin.Context) {
	var list []FK_Struct.Options[uint]
	global.DB.Model(models.BigModelTagModel{}).Select("id as value", "title as label").Scan(&list)
	BY.JSONList(c, len(list), list)
}
