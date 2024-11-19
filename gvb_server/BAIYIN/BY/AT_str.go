package BY

import (
	"fmt"
	"strings"
	"unicode"
)

func (_AT) Trim(s string) string {
	return strings.TrimSpace(s) // 去空
}
func (_AT) Find(s string, str string) bool {
	return strings.Contains(s, str) // Find 是否包含字符串
}
func (_AT) Sp(s string, a ...any) string {
	return fmt.Sprintf(s, a...) // Sp 格式化字符串
}
func (_AT) Split(s string, char string) []string {
	return strings.Split(s, char) // 拆分字符串
}
func (_AT) Replace(s, value, res string) string {
	return strings.ReplaceAll(s, value, res) // Replace 替换字符串
}
func (_AT) Input(obj any, str string) {
	fmt.Printf(str) // 提示输入信息
	fmt.Scan(obj)   // 扫描对象
}

// IsCaps 判断大写字母---字符串|转换次数[-1:转换全部][0:不转换]|头部附加内容
func (_AT) IsCaps(s string, count int, r string) string {
	var n int // 找到次数
	var result string
	for _, ch := range s {
		char := string(ch)
		if unicode.IsUpper(ch) && count != n {
			char = r + strings.ToLower(char) // 小写字母
			n++
		}
		result = result + char
	}
	return result
}

// IsHead 字符串头部是否包含指定字符串
func (_AT) IsHead(s, v string) bool {
	return len(s) >= len(v) && s[0:len(v)] == v
}

// Cut 截取字符串指定长度   str   start   end
func (_AT) Cut(s string, Len ...int) string {
	start := 0
	if len(Len) == 2 {
		start = Len[1]
	}

	abs := []rune(s)
	Leng := len(abs)

	if start > Leng {
		return "" // 起始位置超出字符串长度
	}

	end := start + Len[0]
	if end > Leng {
		end = Leng
	}

	return string(abs[start:end])
}

// RemoveTail 移除最后一个字符
func (_AT) RemoveTail(s string) string {
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}
