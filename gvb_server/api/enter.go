package api

import (
	"gvb_server/api/advertise_api"
	"gvb_server/api/article_api"
	"gvb_server/api/big_model_api"
	"gvb_server/api/chat_api"
	"gvb_server/api/comment_api"
	"gvb_server/api/date_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/message_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/user_api"
)

type GroupApi struct {
	SettingsApi  settings_api.SettingsApi
	ImagesApi    images_api.ImagesApi
	AdvertiseApi advert_api.AdvertApi
	MenuApi      menu_api.MenuApi
	UserApi      user_api.UserApi
	TagApi       tag_api.TagApi
	MessageApi   message_api.MessageApi
	ArticleApi   article_api.ArticleApi
	LikeApi      article_api.LikeApi
	Search       article_api.Search
	CommentsApi  comment_api.CommentApi
	ChatApi      chat_api.ChatApi
	DateApi      date_api.DateApi
	BigModel     big_model_api.BigModelApi
}

var GroupAppApi = new(GroupApi)

const articleLike = "article_like"
const articleLook = "article_look"
const commentCount = "comment_count"
const commentLike = "comment_like"
