package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Driver       string `yaml:"driver"`
	ShowLine     bool   `yaml:"show_line"`      // 显示行号
	LogInConsole bool   `yaml:"log_in_console"` // 显示打印路径
}
