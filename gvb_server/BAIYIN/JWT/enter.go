package JWT

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/INIT/ctype"
)

// Info jwt中payload数据
type Info struct {
	UserName string     `json:"user_name"` // 用户名
	NickName string     `json:"nick_name"` // 昵称
	Role     ctype.Role `json:"role"`      // 用户权限
	UserID   uint       `json:"user_id"`   // 用户id
}

type Claims struct {
	Info
	jwt.StandardClaims
}

const secret = "key" // 秘钥
const issuer = "白音"  // 签发人
const expires = 48   // 过期时间
