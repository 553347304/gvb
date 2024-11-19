package ct

const ArticleLike = "article_like"
const ArticleLook = "article_look"
const CommentCount = "comment_count"
const CommentLike = "comment_like"

const ArticleIndex = "article"
const ArticleParam = `
{
 "settings": { "index":{ "max_result_window": "100000" } },
 "mappings": {
   "properties": {
     "id": { "type": "keyword" },
     "user_id": { "type": "integer" },
     "user_name": { "type": "text" },
     "user_head": { "type": "text" },
     "look_count": { "type": "integer" },
     "comment_count": { "type": "integer" },
     "digg_count": { "type": "integer" },
     "collects_count": { "type": "integer" },
     "banner_id": { "type": "integer" },
     "banner_url": { "type": "text" },
     "title": { "type": "keyword" },
     "abstract": { "type": "text" },
     "content": { "type": "text" },
     "category": { "type": "keyword" },
     "tags": { "type": "keyword" },
     "source": { "type": "text" },
     "link": { "type": "text" },
     "date": { "type": "keyword" },
     "created_at":{
       "type": "date",
       "null_value": "null",
       "format": "[yyyy-MM-dd HH:mm:ss]"
     },
     "updated_at":{
       "type": "date",
       "null_value": "null",
       "format": "[yyyy-MM-dd HH:mm:ss]"
     }
   }
 }
}
`
const SyncArticleIndex = "sync_article"
const SyncArticleParam = `
{
 "settings": { "index":{ "max_result_window": "100000" } },
 "mappings": {
   "properties": {
     "key": { "type": "keyword" },
     "title": { "type": "text" },
     "slug": { "type": "keyword" },
     "content": { "type": "text" }
   }
 }
}
`
const NineArticleIndex = "nine_article"
const NineArticleParam = `
{
 "settings": { "index":{ "max_result_window": "100000" } },
 "mappings": {
   "properties": {
     "id": { "type": "keyword" },
     "user_id": { "type": "integer" },
     "user_name": { "type": "text" },
     "user_head": { "type": "text" },
     "look_count": { "type": "integer" },
     "comment_count": { "type": "integer" },
     "digg_count": { "type": "integer" },
     "collects_count": { "type": "integer" },
     "banner_id": { "type": "integer" },
     "banner_url": { "type": "text" },
     "title": { "type": "keyword" },
     "abstract": { "type": "text" },
     "content": { "type": "text" },
     "category": { "type": "keyword" },
     "tags": { "type": "keyword" },
     "source": { "type": "text" },
     "link": { "type": "text" },
     "date": { "type": "keyword" },
     "created_at":{
       "type": "date",
       "null_value": "null",
       "format": "[yyyy-MM-dd HH:mm:ss]"
     },
     "updated_at":{
       "type": "date",
       "null_value": "null",
       "format": "[yyyy-MM-dd HH:mm:ss]"
     }
   }
 }
}
`
const NineFullIndex = "nine_full"
const NineFullParam = `
{
 "settings": { "index":{ "max_result_window": "100000" } },
 "mappings": {
   "properties": {
     "id": { "type": "keyword" },
     "user_id": { "type": "integer" },
     "user_name": { "type": "text" },
     "user_head": { "type": "text" },
     "look_count": { "type": "integer" },
     "comment_count": { "type": "integer" },
     "digg_count": { "type": "integer" },
     "collects_count": { "type": "integer" },
     "banner_id": { "type": "integer" },
     "banner_url": { "type": "text" },
     "title": { "type": "keyword" },
     "abstract": { "type": "text" },
     "content": { "type": "text" },
     "category": { "type": "keyword" },
     "tags": { "type": "keyword" },
     "source": { "type": "text" },
     "link": { "type": "text" },
     "date": { "type": "keyword" },
     "created_at":{
       "type": "date",
       "null_value": "null",
       "format": "[yyyy-MM-dd HH:mm:ss]"
     },
     "updated_at":{
       "type": "date",
       "null_value": "null",
       "format": "[yyyy-MM-dd HH:mm:ss]"
     }
   }
 }
}
`
