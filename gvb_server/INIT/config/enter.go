package config

type Config struct {
	Logger   Logger   `yaml:"logger"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Email    Email    `yaml:"email"`
	Upload   Upload   `yaml:"upload"`
	System   System   `yaml:"system"`
	BigModel BigModel `yaml:"big_model"`
}
