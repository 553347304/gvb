package models

type ChatModel struct {
	MODEL    `json:","`
	NickName string `gorm:"size:15" json:"nick_name"`
	Avatar   string `gorm:"size:128" json:"avatar"`
	Message  string `gorm:"size:256" json:"message"`
	Addr     string `gorm:"size:64" json:"addr,omit(list)"`
	IsGroup  bool   `json:"is_group"` // 是否是群组消息
	Type     int    `gorm:"size:4" json:"type"`
}
