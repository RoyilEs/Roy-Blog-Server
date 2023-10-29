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

// SettingsInfoUpdateView 更新
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	fmt.Println("1", global.Config)
	//修改
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg(err.Error(), c)
		return
	}
	res.ResultOkWithMsg("修改成功", c)
	fmt.Println("2", global.Config)
}
