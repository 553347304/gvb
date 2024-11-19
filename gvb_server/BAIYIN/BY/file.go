package BY

import (
	"os"
)

func (_File) Read(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		Log("文件不存在")
		return ""
	}
	return string(file)
}

func (_File) Dir(path string) []os.DirEntry {
	dir, err := os.ReadDir(path)
	if err != nil {
		LogE("目录不存在")
		LogE(err)
		return nil
	}
	return dir
}
