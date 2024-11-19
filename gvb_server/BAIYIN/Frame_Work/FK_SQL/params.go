package FK_SQL

import (
	"gorm.io/gorm"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"gvb_server/INIT/global"
)

// 是否开启搜索
func _Param[T any](m T, o FK_Struct.Mysql) *gorm.DB {
	db := global.DB
	for _, preload := range o.Preload {
		db = db.Preload(preload)
	}
	db = db.Model(m) // 查询全部
	if o.Search != "" && o.Key != "" {
		if !mysqlFieldExist(m, o.Search) {
			return db
		}
		switch o.Source {
		case "like":
			db = db.Where(o.Search+" like ?", "%"+o.Key+"%") // 模糊匹配
		default:
			db = db.Where(o.Search+" = ?", o.Key) // 精确匹配
		}
	}

	if o.Sort != "" {
		if !mysqlFieldExist(m, BY.AT.Split(o.Sort, " ")[0]) {
			return db
		}
		db = db.Order(o.Sort)
	}
	return db
}
