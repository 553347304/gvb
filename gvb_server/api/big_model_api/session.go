package big_model_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

type SessionCreateRequest struct {
	RoleId uint `json:"role_id" binding:"required"`
}

func (BigModelApi) SessionCreateView(c *gin.Context) {
	var cr SessionCreateRequest
	if FK.Gin.BindJSON(c, &cr) {
		var role models.BigModelRoleModel
		if !FK_SQL.Is(&role, cr.RoleId) {
			BY.JSON(c, 1000, "会话不存在")
			return
		}
		claims := JWT.Token(c)
		var user models.UserModel
		if !FK_SQL.Is(&user, claims.UserID) {
			BY.JSON(c, 1000, "用户信息错误")
			return
		}
		scope := global.Config.BigModel.Session.SessionScore
		if (user.Scope - scope) <= 0 {
			BY.JSON(c, 1000, "积分不足")
			return
		}

		session := models.BigModelSessionModel{
			UserId: claims.UserID,
			RoleId: cr.RoleId,
			Name:   "新的会话",
		}

		FK_SQL.Create(&session)
		FK_SQL.Update(&user, map[string]interface{}{"scope": gorm.Expr("scope - ?", scope)})
		BY.JSON(c, session.ID, "会话创建成功")
	}
}

type SessionListResponse struct {
	models.MODEL
	UserId      uint   `json:"user_id"`
	NickName    string `json:"nick_name"`
	SessionName string `json:"session_name"`
	RoleName    string `json:"role_name"`
	ChatTotal   int    `json:"chat_total"`
	LastContent string `json:"last_content"`
}

func (BigModelApi) SessionListView(c *gin.Context) {
	var o FK_Struct.Mysql
	if FK.Gin.BindQuery(c, &o) {
		result := FK_SQL.GetList(&models.BigModelSessionModel{}, FK_Struct.Mysql{
			PageInfo: o.PageInfo,
			Preload:  []string{"UserModel", "RoleModel"},
		})

		var list = make([]SessionListResponse, 0)
		for _, m := range result {
			total := len(m.ChatList)
			var content string
			if total != 0 {
				content = m.ChatList[total].Content
			}
			list = append(list, SessionListResponse{
				MODEL:       m.MODEL,
				UserId:      m.UserId,
				NickName:    m.UserModel.NickName,
				SessionName: m.Name,
				RoleName:    m.RoleModel.Name,
				ChatTotal:   total,
				LastContent: content,
			})
		}
		BY.JSONList(c, len(list), list)
	}
}

type SessionUpdateRequest struct {
	SessionId uint   `json:"session_id" binding:"required"` // 会话ID
	Name      string `json:"nick_name" binding:"required"`  // 角色ID
}

func (BigModelApi) SessionUpdateView(c *gin.Context) {
	var cr SessionUpdateRequest
	if FK.Gin.BindJSON(c, &cr) {
		var session models.BigModelSessionModel
		if !FK_SQL.Is(&session, cr.SessionId) {
			BY.JSON(c, 1000, "会话不存在")
			return
		}
		claims := JWT.Token(c)
		if session.ID != claims.UserID {
			BY.JSON(c, 1000, "会话鉴权失败")
			return
		}

		ok := FK_SQL.Update(&session, map[string]interface{}{"name": cr.Name})
		FK.JSON(c, ok, "会话修改")
	}
}

func (BigModelApi) SessionRemoveView(c *gin.Context) {
	var cr FK_Struct.UID
	if FK.Gin.BindURI(c, &cr) {
		var session models.BigModelSessionModel

		if !FK_SQL.Preload("ChatList", &session, cr.ID) {
			BY.JSON(c, 1000, "会话不存在")
			return
		}
		claims := JWT.Token(c)
		if session.UserId != claims.UserID {
			BY.JSON(c, 1000, "会话鉴权失败")
			return
		}

		if len(session.ChatList) > 0 {
			FK_SQL.Delete(session.CreatedAt)
		}
		ok := FK_SQL.Delete(&session)
		FK.JSON(c, ok, "会话删除")
	}
}

func (BigModelApi) SessionAdminRemoveView(c *gin.Context) {
	var id FK_Struct.UID
	if FK.Gin.BindJSON(c, &id) {
		var list []models.BigModelSessionModel
		if !FK_SQL.Preload("ChatList", &list, id.IDList) {
			BY.JSON(c, 1000, "记录不存在")
			return
		}
		for _, i2 := range list {
			FK_SQL.Delete(i2.ChatList)
		}
		ok := FK_SQL.Delete(&list)
		FK.JSON(c, ok, "删除会话", fmt.Sprintf("删除关联对话个数: %d", len(list)))
	}
}
