package date_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/INIT/ct"
	"gvb_server/models"
)

type DataSumResponse struct {
	User         int `json:"user"`           // 用户注册
	UserSameDay  int `json:"user_same_day"`  // 用户当日注册
	LoginSameDay int `json:"login_same_day"` // 用户当日登录
	Message      int `json:"message"`        // 聊天
	ChatMessage  int `json:"chat_message"`   // 群聊
	Article      int `json:"article"`        // 文章
}

func (DateApi) CountView(c *gin.Context) {
	// global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&count)

	BY.JSON(c, "数据总数", DataSumResponse{
		User:         FK_SQL.Find(&models.UserModel{}),
		UserSameDay:  len(FK_SQL.Where(models.UserModel{}, "to_days(created_at)=to_days(now())")),
		LoginSameDay: len(FK_SQL.Where(models.LoginDataModel{}, "to_days(created_at)=to_days(now())")),
		Message:      FK_SQL.Find(&models.MessageModel{}),
		ChatMessage:  len(FK_SQL.Where(models.ChatModel{}, "is_group = ?", true)),
		Article:      FK.ES.Count(ct.ArticleIndex),
	})
}
