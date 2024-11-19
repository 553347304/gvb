package FK

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
)

// JSON 返回前端
func JSON(c *gin.Context, res bool, msg string, data ...any) {
	var result any = res
	if len(data) != 0 {
		result = data[0]
	}
	if !res {
		BY.JSON(c, 1000, msg+"失败", result)
		return
	}
	BY.JSON(c, 0, msg+"成功", result)
}
