package INIT

import (
	"gopkg.in/yaml.v2"
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/config"
	"gvb_server/INIT/global"
	"io/fs"
	"io/ioutil"
)

// Yaml 读取yaml文件的配置
func Yaml() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		BY.LogE("配置文件加载失败: " + err.Error())
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		BY.LogE("config Init Unmarshal: " + err.Error())
	}
	global.Config = c
}

// SetYaml 修改yaml文件
func SetYaml() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		BY.LogE("配置修改失败: " + err.Error())
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		BY.LogE("配置修改失败: " + err.Error())
	}
}
