package BY

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/BAIYIN/System/Time"
	"reflect"
)

func Log(s ...any) {
	fmt.Printf(fmt.Sprintf("[%v] %v ", Time.Now(), getLine(2)))
	fmt.Println(s...)
}
func LogE(s any) {
	fmt.Println(fmt.Sprintf("[%v] %v %v", Time.Now(), getLine(2), color(s, 31)))
	// for i := 2; i < 5; i++ {
	// 	fmt.Println(fmt.Sprintf("%v", getLine(i)))
	// }
}

// JSON 返回前端
func JSON(c *gin.Context, List ...any) {
	code, data, message := param(List...)

	switch Type.T(data) {
	case reflect.String:
		data = data.(string)
	case reflect.Int:
		data = data.(int)
	default:
	}
	c.JSON(200, response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func SSEvent(c *gin.Context, List ...any) {
	code, data, message := param(List...)
	switch Type.T(data) {
	case reflect.String:
		data = data.(string)
	case reflect.Int:
		data = data.(int)
	default:
	}
	byteDate, _ := json.Marshal(response{
		Code:    code,
		Data:    data,
		Message: message,
	})
	c.SSEvent("", string(byteDate))
}

// JSONList 返回前端
func JSONList(c *gin.Context, total int, list any, message ...string) {
	Message := "ok"
	if len(message) != 0 {
		if message[0] != "" {
			Message = message[0]
		}
	}
	c.JSON(200, responseList{
		Code:    0,
		Message: Message,
		Data: dataList{
			Total: total,
			List:  list,
		},
	})
}

// JSONErr 返回前端错误  msg  err
func JSONErr(c *gin.Context, msg string, err error) bool {
	if err != nil {
		JSON(c, 1000, msg, []string{err.Error()})
	}
	return err == nil
}

func JSONRemove(c *gin.Context, err error, ID []uint) {
	if err != nil {
		JSON(c, 1000, "删除失败: "+err.Error())
		return
	}
	JSON(c, fmt.Sprintf("删除数量%d", len(ID)))
}

func LogStruct(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Struct {
		fmt.Printf("结构体:\n")
		subLogStruct(val)
		return
	}
	// 检查传入的是否为切片数组
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			fmt.Printf("第 %d 个结构体:\n", i+1)
			subLogStruct(val.Index(i))
			fmt.Println() // 空行分隔不同结构体的输出
		}
	} else {
		fmt.Println("提供的参数不是切片类型")
	}
}

func subLogStruct(val reflect.Value) {
	typ := val.Type()

	// 检查是否为结构体
	if val.Kind() == reflect.Struct {
		// 遍历结构体字段
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := typ.Field(i)

			// 递归打印子结构体
			if field.Kind() == reflect.Struct {
				fmt.Println("------------------------")
				fmt.Printf("递归打印->结构体字段: %s\n", fieldType.Name)
				subLogStruct(field)
				fmt.Println("------------------------")
				continue
			}

			fmt.Printf("%s: %v\n", fieldType.Name, field.Interface())
		}
	} else {
		fmt.Println("提供的参数不是结构体类型")
	}
}
