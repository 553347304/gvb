package date_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/models"
)

type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
}

func (DateApi) DateView(c *gin.Context) {
	dateList, loginCountList := FK_SQL.DateCount(models.LoginDataModel{}, 7)
	_, signCountList := FK_SQL.DateCount(models.UserModel{}, 7)

	BY.JSON(c, "七日数据总数", DateCountResponse{
		DateList:  dateList,
		LoginData: loginCountList,
		SignData:  signCountList,
	})
}
