package big_model_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

type UserScopeRequest struct {
	Status bool `json:"status"`
}

type Scope struct {
	Enable bool `json:"enable"`
	Scope  int  `json:"scope"`
}

func (BigModelApi) UserIsScopeView(c *gin.Context) {
	claims := JWT.Token(c)
	var userScopeModel models.UserScopeModel
	if FK_SQL.Is(&userScopeModel, "user_id = ? and to_days(created_at)=to_days(now())", claims.UserID) {
		BY.JSON(c, "今日已领取积分了", false)
		return
	}
	scope := global.Config.BigModel.Session.DayScore
	BY.JSON(c, fmt.Sprintf("今日可领取%d积分", scope), true)
}

func (BigModelApi) UserScopeView(c *gin.Context) {
	var cr UserScopeRequest
	if !FK.Gin.BindJSON(c, &cr) {
		return
	}
	claims := JWT.Token(c)
	var userScopeModel models.UserScopeModel
	var response Scope
	if FK_SQL.Is(&userScopeModel, "user_id = ? and to_days(created_at)=to_days(now())", claims.UserID) {
		BY.JSON(c, "今日已领取积分了")
		return
	}
	scope := global.Config.BigModel.Session.DayScore
	response.Enable = true
	response.Scope = scope
	FK_SQL.Create(&models.UserScopeModel{
		UserId: claims.UserID,
		Scope:  scope,
		Status: cr.Status,
	})
	// 加积分
	var user models.UserModel
	if !FK_SQL.Is(&user, claims.UserID) {
		BY.JSON(c, 1000, "用户信息错误")
		return
	}
	result := FK_SQL.Update(&user, map[string]any{"scope": gorm.Expr("scope + ?", scope)})
	BY.JSON(c, "积分领取成功", result)
}
