package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/models"
)

type AutoReplyUpdate struct {
	Id          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`                                         // 规则名称
	Mode        int    `json:"mode" binding:"required,oneof=1 2 3 4" msg:"mode必须为1 2 3 4 其中一个"` // 匹配模型
	Rule        string `json:"rule" binding:"required"`                                         // 匹配规则
	RuleContent string `json:"rule_content" binding:"required"`                                 // 回复内容
}

func (BigModelApi) AutoReplyCreateView(c *gin.Context) {
	var cr AutoReplyUpdate
	if FK.Gin.BindJSON(c, &cr) {
		if FK_SQL.Is(&models.AutoReplyModel{}, "name = ? and id <> ?", cr.Name, cr.Id) {
			BY.JSON(c, 1000, "名字重复")
			return
		}
		R := FK_SQL.Create(&models.AutoReplyModel{
			Name:        cr.Name,
			Mode:        cr.Mode,
			Rule:        cr.Rule,
			RuleContent: cr.RuleContent,
		})
		FK.JSON(c, R, "创建自动回复", R)
	}
}

func (BigModelApi) AutoReplyUpdateView(c *gin.Context) {
	var cr AutoReplyUpdate
	if FK.Gin.BindJSON(c, &cr) {
		var model models.AutoReplyModel
		if !FK_SQL.Is(&model, cr.Id) {
			BY.JSON(c, 1000, "记录不存在")
			return
		}
		if FK_SQL.Is(&model, "name = ? and id <> ?", cr.Name, cr.Id) {
			BY.JSON(c, 1000, "名字重复")
			return
		}
		R := FK_SQL.Update(&model, BY.Map.StructMapDelEmpty(cr))
		FK.JSON(c, R, "更新数据", R)
	}
}

func (BigModelApi) AutoReplyListView(c *gin.Context) {
	var o FK_Struct.Mysql
	if FK.Gin.BindQuery(c, &o) {
		list := FK_SQL.GetList(&models.AutoReplyModel{}, o)
		BY.JSONList(c, len(list), list)
	}
}

func (BigModelApi) AutoReplyRemoveView(c *gin.Context) {
	var id FK_Struct.UID
	if FK.Gin.BindJSON(c, &id) {
		R := FK_SQL.Remove(&[]models.AutoReplyModel{}, id.IDList)
		BY.JSON(c, R.Code, R.Message)
	}
}

func (BigModelApi) AutoReplySearchView(c *gin.Context) {
	reply := models.AutoReplyModel{}
	model := reply.ValidAutoReply("1")
	BY.Log(model)
}
