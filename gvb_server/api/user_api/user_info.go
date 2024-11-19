package user_api

import (
	"github.com/gin-gonic/gin"
	BY "gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/models"
)

// UserInfoView 用户信息
func (UserApi) UserInfoView(c *gin.Context) {
	var cr models.UserModel
	token := JWT.Token(c)
	FK_SQL.Is(&cr, "id = ?", token.UserID)
	BY.JSON(c, cr)
}
