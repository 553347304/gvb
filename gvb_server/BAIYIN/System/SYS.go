package SM

import (
	"os"
	"strings"
)

// Path 获取根目录
func (_SYS) Path() string {
	var dir, _ = os.Getwd()
	return strings.ReplaceAll(dir, "\\", "/")
}

func (_SYS) Exit() {
	os.Exit(0)
}
