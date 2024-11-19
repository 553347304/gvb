package routers

import (
	"gvb_server/BAIYIN/JWT"
	"gvb_server/api"
)

func (router RouterGroup) BigModelRouter() {
	app := api.GroupAppApi.BigModel
	router.GET("big_model/setting", JWT.Auth(), app.ModelSettingView)
	router.PUT("big_model/setting", JWT.Auth(), app.ModelSettingUpdateView)

	router.GET("big_model/scope", JWT.Auth(), app.UserIsScopeView)
	router.POST("big_model/scope", JWT.Auth(), app.UserScopeView)

	router.GET("big_model/auto_reply", JWT.Auth(), app.AutoReplyListView)
	router.POST("big_model/auto_reply", JWT.Auth(), app.AutoReplyCreateView)
	router.PUT("big_model/auto_reply", JWT.Auth(), app.AutoReplyUpdateView)
	router.DELETE("big_model/auto_reply", JWT.Auth(), app.AutoReplyRemoveView)

	// 角色
	{
		router.GET("big_model/tag", JWT.Auth(), app.TagListView)
		router.GET("big_model/tag/options", JWT.Auth(), app.TagOptionsView)
		router.POST("big_model/tag", JWT.Auth(), app.TagCreateView)
		router.PUT("big_model/tag", JWT.Auth(), app.TagUpdateView)
		router.DELETE("big_model/tag", JWT.Auth(), app.TagRemoveView)

		router.GET("big_model/role_sessions", JWT.Auth(), app.RoleSessionListView)

		router.GET("big_model/role", JWT.Auth(), app.RoleListView)
		router.POST("big_model/role", JWT.Auth(), app.RoleCreateView)
		router.PUT("big_model/role", JWT.Auth(), app.RoleUpdateView)
		router.DELETE("big_model/role", JWT.Auth(), app.RoleRemoveView)
		router.GET("big_model/role/:id", JWT.Auth(), app.RoleDetailView)
		router.GET("big_model/role_history", JWT.Auth(), app.RoleUserHistoryListView)

		router.GET("big_model/icon", app.IconView)
	}

	// 会话
	{
		router.GET("big_model/square", JWT.Auth(), app.TagRoleListView)

		router.GET("big_model/session", JWT.Auth(), app.SessionListView)
		router.POST("big_model/session", JWT.Auth(), app.SessionCreateView)
		router.PUT("big_model/session", JWT.Auth(), app.SessionUpdateView)
		router.DELETE("big_model/session/:id", JWT.Auth(), app.SessionRemoveView)
		router.DELETE("big_model/session", JWT.Auth(), app.SessionAdminRemoveView)
	}

	// 对话
	{
		router.GET("big_model/chat_sse", JWT.Auth(), app.ChatCreateView)
		router.GET("big_model/chat", JWT.Auth(), app.ChatListView)
		router.DELETE("big_model/chat/:id", JWT.Auth(), app.ChatRemoveView)
		router.DELETE("big_model/chat", JWT.Auth(), app.ChatAdminRemoveView)
	}

}
