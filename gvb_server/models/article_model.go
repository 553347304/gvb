package models

import (
	"context"
	"github.com/sirupsen/logrus"
	"gvb_server/INIT/global"
)

// ArticleModel 文章表
type ArticleModel struct {
	ID            string   `json:"id"`             // EsID
	UserID        uint     `json:"user_id"`        // 用户id
	UserName      string   `json:"user_name"`      // 用户名字
	UserHead      string   `json:"user_head"`      // 用户头像
	LookCount     int      `json:"look_count"`     // 浏览量
	CommentCount  int      `json:"comment_count"`  // 评论量
	DiggCount     int      `json:"digg_count"`     // 点赞量
	CollectsCount int      `json:"collects_count"` // 收藏量
	BannerID      uint     `json:"banner_id"`      // 文章封面id
	BannerURL     string   `json:"banner_url"`     // 文章封面
	Title         string   `json:"title"`          // 文章标题
	Keyword       string   `json:"keyword"`        // 关键词
	Abstract      string   `json:"abstract"`       // 文章简介
	Content       string   `json:"content"`        // 文章内容
	Category      []string `json:"category"`       // 文章分类
	Tags          []string ` json:"tags"`          // 文章标签
	Source        string   `json:"source"`         // 文章来源
	Link          string   `json:"link"`           // 原文链接
	Date          string   `json:"date"`           // 创建时间
	CreatedAt     string   `json:"created_at"`     // 创建时间
	UpdatedAt     string   `json:"updated_at"`     // 更新时间
}

func (a *ArticleModel) Create() (err error) {
	indexResponse, err := global.ES.Index().Index("article").BodyJson(a).Refresh("true").Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}
