package service_QQ

import (
	"gopkg.in/gomail.v2"
	"gvb_server/INIT/global"
)

type Email struct {
	Receive string `json:"receive" binding:"required" msg:"邮箱不能为空"` // 收件人
	Title   string `json:"title"`                                   // 主题
	Body    string `json:"body"`                                    // 正文
}

// SendEmail 邮件发送  发给谁，主题，正文
func SendEmail(e Email) error {
	var email = global.Config.Email
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(email.User, e.Receive)) // 发件人邮箱.名字
	m.SetHeader("To", e.Receive)                                // 收件人
	m.SetHeader("Subject", e.Title)                             // 主题
	m.SetBody("text/html", e.Body)
	d := gomail.NewDialer(email.Host, email.Port, email.User, email.Password)
	err := d.DialAndSend(m)
	return err
}
