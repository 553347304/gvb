package BY

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func (_Type) BoolInt(b bool) int {
	return map[bool]int{true: 1, false: 0}[b]
}

func (_Type) StructMap(s any) map[string]interface{} {
	if Type.T(s) != reflect.Struct {
		LogE(AT.Sp("不是结构体类型: %v", Type.T(s)))
		return nil
	}
	result := make(map[string]any)
	nameALL := reflect.TypeOf(s)   // 所有字段名
	valueALL := reflect.ValueOf(s) // 所有值
	length := valueALL.NumField()  // 长度

	for i := 0; i < length; i++ {
		// types := BY.Type.T(s, i)                  // 字段类型
		// name := nameALL.Field(i).Name             // 字段名
		field := nameALL.Field(i).Tag.Get("struct")
		switch field {
		case "":
			field = nameALL.Field(i).Tag.Get("json") // Tag名
		case "-":
			continue
		}
		value := valueALL.Field(i).Interface() // 值
		result[field] = value                  // 存入map

		// Log(AT.Sp("%v   %v", field, value))
	}
	// BY.Log(BY.AT.Sp("结构体长度: %v", length))
	return result
}

func (_Type) StructNameToMysqlTableName(m interface{}) string {
	typ := reflect.TypeOf(m)
	if typ.Kind() != reflect.Struct {
		return fmt.Sprintf("不是结构体类型: %s", reflect.String)
	}
	name := typ.Name()

	// Struct convert TableName
	var result strings.Builder
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r)) // 转换为小写字母
		} else {
			result.WriteRune(r) // 小写字母
		}
	}

	return result.String() + "s"
}

func (_Type) String(s any) string {
	return fmt.Sprintf("%v", s)
}
func (_Type) StringInt(T string) int {
	result, err := strconv.Atoi(T)
	if err != nil {
		LogE("转换错误: " + err.Error())
		return -1
	}
	return result
}
func (_Type) StringUint(s string) uint {
	num, err := strconv.ParseUint(s, 10, 0) // 使用 ParseUint 函数
	if err != nil {
		return 0 // 转换错误
	}
	return uint(num) // 将 uint64 to uint
}
func (_Type) DocHtml(s string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	doc.Find("script").Remove()
	return doc.Text()
}
func (_Type) HtmlDoc(s string) string {
	return string(blackfriday.MarkdownCommon([]byte(s)))
}
