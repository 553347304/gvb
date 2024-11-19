package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/service/service_QQ"
)

func (UserApi) SendEmailView(c *gin.Context) {
	var cr service_QQ.Email
	if FK.Gin.BindJSON(c, &cr) {
		code := BY.Random.R(10, 6)
		if cr.Body != "" {
			code = cr.Body
		}
		err := service_QQ.SendEmail(service_QQ.Email{
			Receive: cr.Receive,
			Title:   "验证码",
			Body:    code,
		})
		if BY.JSONErr(c, "发送邮件失败: ", err) {
			BY.JSON(c, "发送成功邮件成功", map[string]string{"code": code})
		}
	}
}
