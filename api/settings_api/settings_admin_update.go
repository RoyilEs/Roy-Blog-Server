package settings_api

import (
	"Goblog/config"
	"Goblog/core"
	"Goblog/global"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsAdminUpdateView(c *gin.Context) {
	var cr config.Admin
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	fmt.Println("1====:", global.Config.Admin)
	//修改
	global.Config.Admin = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg(err.Error(), c)
		return
	}
	res.ResultOkWithMsg("修改成功d=====(￣▽￣*)b", c)
	fmt.Println("2====:", global.Config.Admin)
}
