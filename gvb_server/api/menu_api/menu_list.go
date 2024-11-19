package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}
type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 菜单列表排序
func (MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var ID []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&ID)
	// 查链接表
	var banner []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&banner, "menu_id in ?", ID)
	// 菜单列表
	var list = make([]MenuResponse, 0)
	for _, menu := range menuList {
		list = append(list, MenuResponse{
			MenuModel: menu,
			Banners:   MenuList(menu, banner),
		})
	}

	BY.JSONList(c, len(list), list)

}
