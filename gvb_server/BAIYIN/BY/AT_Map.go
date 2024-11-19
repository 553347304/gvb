package BY

import "reflect"

// IsList 判断列表中是否存在key
func (_Map) IsList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// FirstList 返回第一个值
func (_Map) FirstList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Reverse 反转Map
func (_Map) Reverse(slice []any) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// MapDelEmpty 剔除map空值
func (_Map) MapDelEmpty(list map[string]interface{}) map[string]interface{} {
	var DataMap = map[string]interface{}{}
	for key, v := range list {
		TypeOf := Type.T(v)
		switch TypeOf {
		case reflect.String:
			if v == "" {
				continue
			}
		case reflect.Int, reflect.Uint:
			if AT.Sp("%v", v) == "0" {
				continue
			}
		case reflect.Slice, reflect.Array:
			val := reflect.ValueOf(v)
			if val.Len() == 0 {
				continue
			}
		case reflect.Ptr:
			if reflect.ValueOf(v).IsNil() {
				continue // 空指针
			}
		default:
		}
		DataMap[key] = v
	}
	return DataMap
}

// StructMapDelEmpty 结构体转Map并剔除空值
func (by _Map) StructMapDelEmpty(s any) map[string]interface{} {
	_map := Type.StructMap(s)
	return by.MapDelEmpty(_map)
}
