package config

type System struct {
	Port  string `yaml:"port"`
	Mysql string `yaml:"mysql"`
	Es    string `yaml:"es"`
	Redis string `yaml:"redis"`
	Log   string `yaml:"log"`
}
