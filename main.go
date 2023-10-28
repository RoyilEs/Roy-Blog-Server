package main

import (
	"Goblog/core"
	"Goblog/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Info("初始化日志成功")
	global.Log.Error("初始化日志成功")
	global.Log.Warning("初始化日志成功")
	//连接数据库
	global.DB = core.InitGorm()

}
