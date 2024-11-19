package models

// CommentModel 评论表
type CommentModel struct {
	MODEL            // 父级评论
	ParentId  uint   `json:"parent_id"`  // 父评论id
	ArticleID string `json:"article_id"` // 文章id
	User      uint   `json:"user"`       // 评论的用户
	Content   string `json:"content"`    // 评论内容
}
