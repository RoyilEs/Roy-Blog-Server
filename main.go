package main

import (
	"Goblog/core"
	_ "Goblog/docs"
	"Goblog/flag"
	"Goblog/global"
	"Goblog/routers"
)

// @title Goblog API文档
// @version 1.0
// @description Goblog API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("blog is running at: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
