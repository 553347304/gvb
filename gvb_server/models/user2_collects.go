package models

import "time"

// UserCollectModel 自定义第三张收藏表
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"primaryKey"`
	CreatedAt time.Time
}
