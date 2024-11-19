package models

import (
	"gorm.io/gorm"
	"gvb_server/INIT/ctype"
	"os"
)

// BannerModel banner表
type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片类型
}

// BeforeDelete 勾子函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 删除本地图片
	if b.ImageType == ctype.Local {
		err = os.Remove(b.Path)
	}
	return
}
