package FK

import (
	"gvb_server/BAIYIN/BY"
)

// Email 邮箱脱敏
func (_Desensitization) Email(email string) string {
	eList := BY.AT.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

// Tel 手机号脱敏
func (_Desensitization) Tel(tel string) string {
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
