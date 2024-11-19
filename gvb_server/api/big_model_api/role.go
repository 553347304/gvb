package big_model_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/BY"
	FK "gvb_server/BAIYIN/Frame_Work"
	"gvb_server/BAIYIN/Frame_Work/FK_SQL"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/JWT"
	"gvb_server/INIT/global"
	"gvb_server/models"
	"strconv"
	"time"
)

type RoleRequest struct {
	Id        uint   `json:"id"`
	Name      string `binding:"required" json:"name"`
	Enable    bool   `json:"enable"`
	Icon      string `json:"icon"` // 系统默认   自己上传
	Abstract  string `binding:"required" json:"abstract"`
	Scope     int    `json:"scope"`                       // 消耗积分
	Prologue  string `binding:"required" json:"prologue"` // 开场白
	Prompt    string `binding:"required" json:"prompt"`   // 设定词
	AutoReply bool   `json:"auto_reply"`
	TagList   []uint `json:"tag_list" struct:"-"` // 标签ID列表
}

func isTag(c *gin.Context, cr RoleRequest) ([]models.BigModelTagModel, int) {
	// 先找标签
	var tags []models.BigModelTagModel
	global.DB.Find(&tags, cr.TagList)
	if len(cr.TagList) != len(tags) {
		BY.JSON(c, 1000, "标签不一致")
		return tags, -1
	}
	return tags, len(tags)
}

func (BigModelApi) RoleCreateView(c *gin.Context) {
	var cr RoleRequest
	if FK.Gin.BindJSON(c, &cr) {
		if FK_SQL.Is(&models.BigModelRoleModel{}, "name = ?", cr.Name) {
			BY.JSON(c, 1000, "名字重复")
			return
		}
		tags, total := isTag(c, cr)
		if total == -1 {
			return
		}
		R := FK_SQL.Create(&models.BigModelRoleModel{
			Name:      cr.Name,
			Enable:    cr.Enable,
			Icon:      cr.Icon,
			Abstract:  cr.Abstract,
			Scope:     cr.Scope,
			Prologue:  cr.Prologue,
			Prompt:    cr.Prompt,
			AutoReply: cr.AutoReply,
			Tags:      tags, // 会自动关联表
		})
		FK.JSON(c, R, "角色添加", R)
	}
}

func (BigModelApi) RoleUpdateView(c *gin.Context) {
	var cr RoleRequest
	if FK.Gin.BindJSON(c, &cr) {
		var model models.BigModelRoleModel
		err := global.DB.Preload("Tags").Take(&model, cr.Id).Error
		if err != nil {
			BY.JSON(c, 1000, "记录不存在")
			return
		}
		tags, total := isTag(c, cr)
		if total == -1 {
			return
		}
		err = FK_SQL.Association(&model, "Tags").Replace(tags) // 替换
		if err != nil {
			BY.JSON(c, 1000, "更新失败", err)
			return
		}
		R := FK_SQL.Update(&model, BY.Map.StructMapDelEmpty(cr))
		FK.JSON(c, R, "更新数据", R)
	}
}

func (BigModelApi) RoleListView(c *gin.Context) {
	var o FK_Struct.Mysql
	if FK.Gin.BindQuery(c, &o) {
		list := FK_SQL.GetList(&models.BigModelRoleModel{}, FK_Struct.Mysql{
			PageInfo: o.PageInfo,
			Preload:  []string{"Tags"},
		},
		)
		BY.JSONList(c, len(list), list)
	}
}

func (BigModelApi) RoleRemoveView(c *gin.Context) {
	var id FK_Struct.UID
	if FK.Gin.BindJSON(c, &id) {
		var list []models.BigModelRoleModel
		count := global.DB.Preload("Tags").Find(&list, id.IDList).RowsAffected
		if count == 0 {
			BY.JSON(c, 1000, "记录不存在")
			return
		}
		for _, i2 := range list {
			err := FK_SQL.Association(&i2, "Tags").Delete(i2.Tags) // 删除关联关系
			if err != nil {
				BY.LogE("Tags删除失败")
				return
			}
		}
		err := global.DB.Delete(&list).Error
		FK.JSON(c, err == nil, "删除角色", err == nil)
	}
}

type RoleSessionRequest struct {
	FK_Struct.PageInfo
	RoleId uint `form:"role_id" binding:"required"`
}
type RoleSessionResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (BigModelApi) RoleSessionListView(c *gin.Context) {
	var cr RoleSessionRequest
	if FK.Gin.BindQuery(c, &cr) {
		claims := JWT.Token(c)
		result := FK_SQL.GetList(models.BigModelSessionModel{UserId: claims.UserID, RoleId: cr.RoleId}, FK_Struct.Mysql{
			PageInfo: FK_Struct.PageInfo{
				Search: "role_id",
				Key:    strconv.Itoa(int(cr.RoleId)),
				Sort:   "created_at desc",
				Page:   cr.Page,
				Limit:  cr.Limit,
			},
		})
		var list = make([]RoleSessionResponse, 0)
		for _, m := range result {
			list = append(list, RoleSessionResponse{
				Id:        m.ID,
				Name:      m.Name,
				CreatedAt: m.CreatedAt,
			})
		}
		BY.JSONList(c, len(list), list)
	}
}
