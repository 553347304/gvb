package comment_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_ES"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/ct"
	"gvb_server/models"
)

type CommentRequest struct {
	ParentId  uint   `json:"parent_id"`
	ArticleID string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content   string `json:"content" binding:"required" msg:"请输入评论内容"`
}

func (CommentApi) CreateView(c *gin.Context) {
	var cr CommentRequest
	if FK.Gin.BindJSON(c, &cr) {
		claims := JWT.Token(c)
		is := FK_ES.Get(ct.ArticleIndex, FK_ES.ID, cr.ArticleID, "")
		if is == "" {
			BY.JSON(c, "文章不存在")
			return
		}

		// 添加评论
		R := FK_SQL.Create(&models.CommentModel{
			ParentId:  cr.ParentId,
			ArticleID: cr.ArticleID,
			User:      claims.UserID,
			Content:   cr.Content,
		})
		FK.Redis.HAdd(ct.CommentCount, cr.ArticleID, 1)
		BY.JSON(c, "添加文章评论", R)
	}
}
