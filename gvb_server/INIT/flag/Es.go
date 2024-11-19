package flag

import (
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/INIT/ct"
)

func Es(s string) {
	FK.ES.DelIndex(ct.NineArticleIndex)
	FK.ES.DelIndex(ct.NineFullIndex)
	FK.ES.DelIndex(ct.ArticleIndex)
	FK.ES.DelIndex(ct.SyncArticleIndex)
	if s != "d" {
		FK.ES.CreateIndex(ct.NineArticleIndex, ct.NineArticleParam)
		FK.ES.CreateIndex(ct.NineFullIndex, ct.NineFullParam)
		FK.ES.CreateIndex(ct.ArticleIndex, ct.ArticleParam)
		FK.ES.CreateIndex(ct.SyncArticleIndex, ct.SyncArticleParam)
	}
}
