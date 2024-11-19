package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_ES"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/BAIYIN/System/Time"
	"gvb_server/INIT/ct"
	"gvb_server/models"
	"gvb_server/service/ser_es"
)

type ArticleRequest struct {
	BannerID  uint     `json:"banner_id"`  // 文章封面id
	BannerURL string   `json:"banner_url"` // 文章封面
	Title     string   `json:"title"`      // 文章标题
	Abstract  string   `json:"abstract"`   // 文章简介
	Content   string   `json:"content"`    // 文章内容
	Category  []string `json:"category"`   // 文章分类
	Tags      []string ` json:"tags"`      // 文章标签
	Source    string   `json:"source"`     // 文章来源
	Link      string   `json:"link"`       // 原文链接
}

func (ArticleApi) CreateView(c *gin.Context) {
	var cr ArticleRequest
	if FK.Gin.BindJSON(c, &cr) {
		user := JWT.Token(c)
		// 没写Abstract截取标题指定个数字符串
		if cr.Abstract == "" {
			html := BY.Type.DocHtml(cr.Content)
			Content := BY.Type.HtmlDoc(html)
			cr.Abstract = BY.AT.Cut(Content, 100)
		}

		title := FK_ES.Get(ct.ArticleIndex, "title", cr.Title, "title")
		if title != "" {
			BY.JSON(c, msg+"已存在")
			return
		}
		BY.Log(title)
		BannerURL := FK_SQL.Get(models.BannerModel{}, "id", cr.BannerID, "path")
		UserHead := FK_SQL.Get(models.UserModel{}, "avatar", user.UserID, "avatar")
		// 入库
		article := models.ArticleModel{
			Title:     cr.Title,
			Keyword:   cr.Title,
			Abstract:  cr.Abstract,
			Content:   cr.Content,
			Category:  cr.Category,
			Tags:      cr.Tags,
			Source:    cr.Source,
			Link:      cr.Link,
			Date:      Time.Now(),
			CreatedAt: Time.Now(),
			UpdatedAt: Time.Now(),
			UserID:    user.UserID,
			UserName:  user.UserName,
			UserHead:  UserHead,
			BannerID:  cr.BannerID,
			BannerURL: BannerURL,
		}

		id := FK_ES.Create(ct.ArticleIndex, article)
		FK.ES.Update(ct.ArticleIndex, id, map[string]any{"id": id})
		FK_ES.Create(ct.SyncArticleIndex, ser_es.GetHtml(id, cr.Title, cr.Content))

		BY.JSON(c, msg+"创建成功", id)
	}
}
