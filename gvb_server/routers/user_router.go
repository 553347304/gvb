package routers

import (
	"gvb_server/BAIYIN/JWT"
	"gvb_server/api"
)

func (router RouterGroup) UserRouter() {
	app := api.GroupAppApi.UserApi
	const index = "user"
	router.POST("email_login", app.EmailLoginView)
	router.POST("send_email", JWT.Auth(), app.SendEmailView)
	router.GET("user_info", JWT.Auth(), app.UserInfoView)
	router.PUT("user_role", JWT.Auth(), app.UserUpdateRoleView)
	router.PUT("user_password", JWT.Auth(), app.UserUpdatePasswordView)
	router.POST("user_logout", JWT.Auth(), app.UserLogoutView)

	router.POST(index, app.UserCreateView)
	router.DELETE(index, JWT.Auth(), app.UserRemoveView)
	router.GET(index, JWT.Auth(), app.UserListView)

}
