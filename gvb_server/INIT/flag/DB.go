package flag

import (
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

func MigrationTable() {
	var err error
	// global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectsModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.UserCollectModel{},
			&models.CommentModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
			&models.ChatModel{},
			&models.LoginDataModel{},
			&models.UserScopeModel{},
			&models.AutoReplyModel{},
			&models.BigModelTagModel{},
			&models.BigModelRoleModel{},
			&models.BigModelChatModel{},
			&models.BigModelSessionModel{},
		)
	if err != nil {
		BY.LogE("[生成数据库表结构失败]")
		return
	}
	BY.Log("[生成数据库表结构成功]")
}
