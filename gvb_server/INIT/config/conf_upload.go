package config

type Upload struct {
	Size int    `yaml:"size" json:"size"` // 文件大小限制
	Path string `yaml:"path" json:"path"` // 文件保存路径
}
