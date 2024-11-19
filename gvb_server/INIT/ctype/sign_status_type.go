package ctype

import "encoding/json"

type LoginMethod int

const (
	SignQQ    LoginMethod = 1 // service_QQ
	SignGitee LoginMethod = 2 // Gitee
	SignEmail LoginMethod = 3 // 邮箱
)

func (s LoginMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s LoginMethod) String() string {
	switch s {
	case SignQQ:
		return "service_QQ"
	case SignGitee:
		return "Gitee"
	case SignEmail:
		return "邮箱"
	}
	return "其他"
}
