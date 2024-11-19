package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

type TagRequest struct {
	Title string `json:"title"`
}

func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	if FK.Gin.BindJSON(c, &cr) {
		if FK_SQL.Is(&models.TagModel{}, "title = ?", cr.Title) {
			BY.JSON(c, msg+"已存在")
			return
		}
		res := FK_SQL.Create(&models.TagModel{
			Title: cr.Title,
		})
		BY.JSON(c, "创建"+msg, res)
	}
}
