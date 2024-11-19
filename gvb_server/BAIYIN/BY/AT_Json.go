package BY

import (
	"encoding/json"
	"github.com/liu-cn/json-filter/filter"
)

// Select 直接返回过滤后的数据结构，它可以被json.Marshal解析，直接打印会以过滤后的json字符串展示
func (_Json) Select(s string, el interface{}) interface{} {
	return filter.Select(s, el)
}

// Omit 排除指定字段   json:",omit(name)"`
func (_Json) Omit(o string, cr interface{}) interface{} {
	return filter.Omit(o, cr)
}

// Unmarshal byte转json
func (_Json) Unmarshal(data []byte, obj any) bool {
	err := json.Unmarshal(data, obj)
	if err != nil {
		Log("json转换失败: " + err.Error())
		return false
	}
	return true
}
func (_Json) Marshal(Map any) []byte {
	jsonData, _ := json.Marshal(Map)
	return jsonData
}
