package models

import (
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
	"regexp"
	"strings"
)

type AutoReplyModel struct {
	MODEL
	Name        string `gorm:"size:32" json:"name"`           // 规则名称
	Mode        int    `json:"mode"`                          // 匹配模型
	Rule        string `gorm:"size:64" json:"rule"`           // 匹配规则
	RuleContent string `gorm:"size:1024" json:"rule_content"` // 回复内容
}

// ValidAutoReply 是否命中自动回复
func (AutoReplyModel) ValidAutoReply(content string) *AutoReplyModel {
	var list []AutoReplyModel
	global.DB.Find(&list)
	for _, model := range list {

		switch model.Mode {
		case 1:
			// 精确
			if model.Rule == content {
				return &model
			}
		case 2:
			// 包含
			if strings.Contains(model.Rule, content) {
				return &model
			}
		case 3:
			// 前缀
			if strings.HasPrefix(model.Rule, content) {
				return &model
			}
		case 4:
			// 正则
			regex, err := regexp.Compile(model.Rule)
			if err != nil {
				BY.LogE("正则语法错误")
				return nil
			}
			if regex.MatchString(content) {
				return &model
			}
		}
	}
	return nil
}
