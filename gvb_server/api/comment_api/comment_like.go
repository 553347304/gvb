package comment_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/ct"
	"gvb_server/models"
)

func (CommentApi) LikeView(c *gin.Context) {
	var cr FK_Struct.ID
	if FK.Gin.BindURI(c, &cr) {
		var model models.CommentModel
		if !FK_SQL.Is(&model, cr.ID) {
			BY.JSON(c, "评论不存在")
			return
		}
		FK.Redis.HAdd(ct.CommentLike, cr.ID, 1)
		BY.JSON(c, "评论点赞成功")
	}
}
