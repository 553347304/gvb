package FK_SQL

import (
	"fmt"
	"gorm.io/gorm"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/BAIYIN/System/Time"
	"gvb_server/INIT/global"
	"strings"
)

// func Take[T any](m *T, dest interface{}, arms ...interface{}) bool {
// 	return global.DB.Take(m, dest, arms).Error == nil
// }

// Get 根据x获取一个字段的值
func Get[T any](m T, search string, value interface{}, res string) string {
	var result string
	global.DB.Model(m).Where(search+" = ?", value).Select(res).Scan(&result)
	return result
}

func Is[T any](m *T, q interface{}, args ...interface{}) bool {
	var total int64
	global.DB.Model(m).Where(q, args...).Count(&total).Find(&m)
	return total != 0
}

func Find[T any](m *T, opts ...interface{}) int {
	var total int64
	global.DB.Model(m).Count(&total).Find(&m, opts...)
	return int(total)
}

func Where[T any](m T, q interface{}, args ...interface{}) []T {
	var list = make([]T, 0)
	global.DB.Model(m).Where(q, args...).Find(&list)
	return list
}

func Group[T any](m any, o FK_Struct.Mysql, r *[]T, s string) {
	db := global.DB.Model(m)
	if len(o.Where) > 0 {
		db = db.Where(o.Where[0], o.Where[1:]...)
	}
	db.Group(s).Select(s).Scan(&r)
}

func FindGroup[T any](m any, o FK_Struct.Mysql, r *[]T, s string) {
	o.Param()
	o.Sort = "" // 不支持排序
	DB := _Param(m, o)

	var group string
	list := strings.Split(s, ",")
	for _, f := range list {
		if strings.Contains(f, "count(") {
			continue
		}
		field := strings.Split(f, " as ")
		group = group + field[1] + ","
	}
	group = BY.AT.RemoveTail(group) // 移除尾部,

	DB = DB.Group(group).Select(s) // +
	DB.Limit(o.Limit).Offset(o.Page).Find(&r)
}

// GetList 列表查询
func GetList[T any](m T, o FK_Struct.Mysql) []T {
	var list = make([]T, 0)
	o.Param()
	DB := _Param(m, o)
	DB.Limit(o.Limit).Offset(o.Page).Find(&list) // 分页查
	return list
}

// Remove 根据ID批量删除
func Remove[T any](m *[]T, ID []uint) TypeResponseRemove {
	err := global.DB.Delete(m, ID).Error
	if err != nil {
		return TypeResponseRemove{Code: 1000, Message: "删除失败: " + err.Error()}
	}
	return TypeResponseRemove{Message: fmt.Sprintf("删除数量%d", len(ID))}
}

func Delete(obj interface{}, value ...interface{}) bool {
	return global.DB.Delete(obj, value...).Error == nil
}

func Create[T any](m *T) bool {
	return global.DB.Create(&m).Error == nil
}

func Update[T any](m *T, Map map[string]interface{}) bool {
	return global.DB.Model(m).Updates(Map).Error == nil
}

func Association[T any](m *T, a string) *gorm.Association {
	return global.DB.Model(&m).Association(a)
}
func Preload[T any](p string, m *T, opts ...interface{}) bool {
	total := global.DB.Preload(p).Find(&m, opts...).RowsAffected
	return total != 0
}

func DateCount[T any](m T, d int) (date []string, total []int) {
	var result []struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}
	global.DB.Model(m).Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= created_at", d)).
		Select("date_format(created_at,'%Y-%m-%d') as date", "count(id) as count").
		Group("date").Scan(&result)

	var Map = make(map[string]int)
	for _, i := range result {
		Map[i.Date] = i.Count
	}

	for i := -d; i != 0; i++ {
		day := Time.NowDay(i)
		date = append(date, day)
		total = append(total, Map[day])
	}
	return
}
