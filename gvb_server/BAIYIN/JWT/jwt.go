package JWT

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/BAIYIN/BY"
	"time"
)

// GenToken 加密 Token
func GenToken(user Info) (string, error) {
	MySecret := []byte(secret) // 秘钥
	claim := Claims{user, jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(expires))), // 过期时间
		Issuer:    issuer,                                                     // 签发人
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret) // 创建Token
}

// ParseToken 解析 Token
func ParseToken(tokenStr string) *Claims {
	if tokenStr == "" {
		BY.LogE("未携带token")
		return nil
	}
	MySecret := []byte(secret)
	token, _ := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		return MySecret, nil
	})

	if token != nil {
		// 验证token
		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			return claims
		}
	}
	BY.LogE("token错误")
	return nil
}
