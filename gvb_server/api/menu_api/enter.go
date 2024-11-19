package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

type MenuApi struct{}

type MenuModel struct {
	Title        string `json:"title" binding:"required" msg:"请输入菜单名称"`
	Path         string `json:"path" binding:"required" msg:"请输入菜单路径"`
	Slogan       string `json:"slogan"`
	Abstract     string `json:"abstract"`
	AbstractTime int    `json:"abstract_time"`                         // 切换时间
	Banner       []uint `json:"banner"`                                // 背景图
	BannerTime   int    `json:"banner_time"`                           // 切换时间
	Sort         int    `json:"sort" binding:"required" msg:"请输入菜单序号"` // 菜单序号
}

func create(cr MenuModel) bool {
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	FK_SQL.Create(&menuModel) // 创建菜单
	var bannerList []models.MenuBannerModel
	for i, id := range cr.Banner {
		bannerList = append(bannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: id,
			Sort:     i,
		})
	}
	R := FK_SQL.Create(&bannerList)
	return R
}

func Remove(c *gin.Context, id []uint) bool {
	var List []models.MenuModel
	is := FK_SQL.Preload("Banners", &List, id)
	if !is {
		BY.JSON(c, 1000, "ID不存在")
		return false
	}

	// 事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		for _, i2 := range List {
			err := FK_SQL.Association(&i2, "Banners").Delete(i2.Banners) // 删除关联关系
			if err != nil {
				return fmt.Errorf("banners删除失败:" + err.Error())
			}
			if !FK_SQL.Delete(&i2) {
				return fmt.Errorf("删除失败")
			}
		}
		return nil
	})
	if err != nil {
		BY.JSON(c, "删除菜单失败", err)
		return false
	}
	return true
}

func MenuList(model models.MenuModel, banner []models.MenuBannerModel) []Banner {
	var banners = make([]Banner, 0) // 默认值null改成[]
	for _, b := range banner {
		if model.ID != b.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   b.BannerID,
			Path: b.BannerModel.Path,
		})
	}
	return banners
}
