package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"` // ID列表
}

/* 表结构说明
advert_model.go          广告表
user_collect_model.go    用户收藏文章表
menu_banner_model.go     菜单banner表
menu_model.go            菜单表
File.go
user_model.go            用户表
banner_model.go          banner表
comment_model.go         评论表
tag_model.go             标签表
login_data_model.go      登录信息表
article_model.go         文章表
fade_back_model.go       用户反馈表
message_model.go         消息表
*/
