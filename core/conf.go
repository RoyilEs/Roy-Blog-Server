package core

import (
	"Goblog/config"
	"Goblog/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// InitConf conf初始化读取配置文件
func InitConf() {
	const ConfigFile = "application.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %v", err))
	}
	//读取配置文件
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("yaml.Unmarshal error:%v", err)
	}
	log.Println("config yamlFile InitConf success")
	//设置到global-config
	global.Config = c
}
