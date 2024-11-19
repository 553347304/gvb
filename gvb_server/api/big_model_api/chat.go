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
	"gvb_server/INIT/ctype"
	"gvb_server/INIT/global"
	"gvb_server/models"
	"gvb_server/service/big_model_ser"
	"io"
	"strconv"
)

type ChatListRequest struct {
	SessionId uint `json:"session_id" form:"session_id"`
	FK_Struct.PageInfo
}
type ChatListResponse struct {
	models.MODEL
	UserContent string `json:"user_content"`
	UserAvatar  string `json:"user_avatar"`
	BotContent  string `json:"bot_content"`
	BotAvatar   string `json:"bot_avatar"`
	Status      bool   `json:"status"`
}

func auth[T any](c *gin.Context, m *T, id uint) bool {
	if !FK_SQL.Is(&m, id) {
		BY.JSON(c, 1000, "ID不存在")
		return false
	}
	claims := JWT.Token(c)
	if claims.Role != ctype.PermissionAdmin {
		BY.JSON(c, 1000, "鉴权失败")
		return false
	}
	return true
}

type ChatCreateRequest struct {
	SessionId uint   `form:"session_id" binding:"required"`
	Content   string `form:"content" binding:"required"` // 对话内容
}

func (BigModelApi) ChatCreateView(c *gin.Context) {

	var cr ChatCreateRequest
	if FK.Gin.BindQuery(c, &cr) {
		var session models.BigModelSessionModel
		if !FK_SQL.Is(&session, cr.SessionId) {
			BY.JSON(c, 1000, "会话不存在")
			return
		}

		claims := JWT.Token(c)
		if claims == nil {
			return
		}

		if session.UserId != claims.UserID {
			BY.SSEvent(c, 1000, "会话鉴权错误")
			return
		}

		var user models.UserModel
		if !FK_SQL.Is(&user, claims.UserID) {
			BY.JSON(c, 1000, "用户信息错误")
			return
		}
		scope := global.Config.BigModel.Session.ChatScore
		if (user.Scope - scope) <= 0 {
			BY.SSEvent(c, 1000, "积分不足")
			return
		}

		channel, err := big_model_ser.Send(cr.SessionId, cr.Content)
		if err != nil {
			BY.Log(err)
			return
		}
		var botContent string
		c.Stream(func(w io.Writer) bool {
			s, ok := <-channel
			if ok {
				BY.SSEvent(c, "ok", s)
				botContent += s
				return true
			}
			return false
		})

		chatModel := models.BigModelChatModel{
			SessionId:  cr.SessionId,
			Status:     true,
			Content:    cr.Content,
			BotContent: botContent,
			UserId:     session.RoleId,
			RoleId:     claims.UserID,
		}
		ok := FK_SQL.Create(&chatModel)
		if !ok {
			BY.SSEvent(c, 1000, "对话失败")
			return
		}

		// 对话成功 扣用户积分
		FK_SQL.Update(&user, map[string]interface{}{"scope": gorm.Expr("scope - ?", scope)})
		BY.JSON(c, chatModel.ID, "对话创建成功")
	}
}

func (BigModelApi) ChatListView(c *gin.Context) {
	var cr ChatListRequest
	if FK.Gin.BindQuery(c, &cr) {
		if !auth(c, &models.BigModelSessionModel{}, cr.SessionId) {
			return
		}
		result := FK_SQL.GetList(models.BigModelChatModel{}, FK_Struct.Mysql{
			PageInfo: FK_Struct.PageInfo{
				Search: "session_id",
				Key:    strconv.Itoa(int(cr.SessionId)),
				Page:   cr.Page,
				Limit:  cr.Limit,
				Sort:   "created_at asc",
				Source: cr.Source,
			},
			Preload: []string{"UserModel", "RoleModel"},
		})
		var list = make([]ChatListResponse, 0)
		for _, m := range result {
			list = append(list, ChatListResponse{
				MODEL:       m.MODEL,
				UserContent: m.Content,
				UserAvatar:  m.UserModel.Avatar,
				BotContent:  m.BotContent,
				BotAvatar:   m.RoleModel.Icon,
				Status:      m.Status,
			})
		}
		BY.JSONList(c, len(list), list)
	}
}

func (BigModelApi) ChatRemoveView(c *gin.Context) {
	var id FK_Struct.UID
	if FK.Gin.BindURI(c, &id) {
		var chat models.BigModelChatModel
		claims := JWT.Token(c)
		if !auth(c, &chat, id.ID) {
			return
		}
		if chat.UserId != claims.UserID {
			BY.JSON(c, 1000, "对话鉴权失败")
			return
		}
		is := FK_SQL.Delete(&chat)
		FK.JSON(c, is, "对话")
	}
}

func (BigModelApi) ChatAdminRemoveView(c *gin.Context) {
	var id FK_Struct.UID
	if FK.Gin.BindJSON(c, &id) {
		var list []models.BigModelChatModel
		if FK_SQL.Find(&list, id.IDList) == 0 {
			BY.JSON(c, 1000, "记录不存在")
			return
		}
		ok := FK_SQL.Delete(&list)
		FK.JSON(c, ok, "删除对话", fmt.Sprintf("删除对话个数: %d", len(list)))
	}
}
