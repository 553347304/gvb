package models

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `json:"title" binding:"required" msg:"请输入标题"`       // 显示的标题
	Href   string `json:"href" binding:"required"`                    // 跳转链接
	Images string `json:"images" binding:"required,uri" msg:"非法图片地址"` // 图片
	IsShow bool   `json:"is_show"`                                    // 是否展示
}
