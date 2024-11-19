package BY

import (
	"fmt"
	SM "gvb_server/BAIYIN/System"
	"reflect"
	"runtime"
	"strings"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type dataList struct {
	Total int `json:"total"`
	List  any `json:"list"`
}
type responseList struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    dataList `json:"data"`
}

var selfPath = SM.SYS.Path() // 工作目录

var codeMap = map[int]string{
	1000: "错误",
}

// color 日志颜色
func color(s any, c int) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}

// getLine 获取路径行号
func getLine(skip int) string {
	_, Path, line, _ := runtime.Caller(skip) // 获取调用者---级别
	Remove := strings.ReplaceAll(Path, selfPath, "")
	path := color(fmt.Sprintf("[%v:%v]", Remove, line), 94)
	return path
}

// param 根据类型返回前端数据
func param(List ...any) (code int, data any, msg string) {
	var Len = len(List)
	data = map[string]any{}
	var isString = false
	var isInt = false
	for i := 0; i < Len; i++ {
		switch Type.T(List[i]) {
		case reflect.String:
			if isString {
				data = List[i].(any)
				continue
			}
			msg = List[i].(string)
			isString = true
		case reflect.Int:
			if isInt {
				data = List[i].(any)
				continue
			}
			code = List[i].(int)
			isInt = true
		default:
			data = List[i].(any)
		}
	}
	if msg == "" {
		msg = codeMap[code]
		if msg == "" {
			msg = "ok"
		}
	}
	return
}
