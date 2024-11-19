package BY

import (
	"reflect"
)

// T 获取类型
func (_Type) T(T ...any) reflect.Kind {
	TypeOf := reflect.TypeOf(T[0])
	// 获取子类型
	if len(T) == 2 {
		TypeOf = TypeOf.Field(T[1].(int)).Type
	}
	return TypeOf.Kind()
}
