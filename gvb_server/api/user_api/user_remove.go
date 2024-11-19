package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/INIT/global"
	"gvb_server/models"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	if FK.Gin.BindJSON(c, &cr) {
		var model []models.UserModel
		if FK_SQL.Is(c, &model, cr.IDList) {
			// 事务
			err := global.DB.Transaction(func(tx *gorm.DB) error {
				if !FK_SQL.Delete(&model) {
					return fmt.Errorf("无法删除有数据的用户")
				}
				return nil
			})
			if err != nil {
				BY.JSON(c, 1000, err.Error())
				return
			}
			BY.JSON(c, BY.AT.Sp("删除个数%d", len(model)))
		}
	}
}
