package main

import (
	"Goblog/core"
	"Goblog/global"
	"fmt"
)

func main() {
	//读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
}
