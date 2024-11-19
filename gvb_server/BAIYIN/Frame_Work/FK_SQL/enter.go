package FK_SQL

import (
	"fmt"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
)

type TypeResponseRemove struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func mysqlFieldExist(m interface{}, fieldName string) bool {
	is := global.DB.Migrator().HasColumn(m, fieldName)
	if !is {
		tableName := BY.Type.StructNameToMysqlTableName(m)
		BY.LogE(fmt.Sprintf("'%s' 表中不存在 '%s' 字段", tableName, fieldName))
	}
	return is
}
