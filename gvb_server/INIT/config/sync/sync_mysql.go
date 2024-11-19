package sync

import (
	"gorm.io/gorm"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/INIT/ct"
	"gvb_server/models"
)

func Mysql() {
	commentLikeInfo := FK.Redis.HGetAll(ct.CommentLike)
	for key, count := range commentLikeInfo {
		var comment models.CommentModel
		if !FK_SQL.Is(&comment, key) {
			continue
		}

		FK_SQL.Update(&comment, map[string]interface{}{"digg_count": gorm.Expr("digg_count + ?", count)})
		BY.Log(comment.Content)
	}
	FK.Redis.Del(ct.CommentLike) // 删除
	BY.Log("Mysql同步成功")
}
