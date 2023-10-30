package settings_api

import (
	"Goblog/global"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 视图
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.ResultOkWithData(global.Config.SiteInfo, c)
	case "qiniu":
		res.ResultOkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.ResultOkWithData(global.Config.Jwt, c)
	default:
		res.ResultFailWithMsg("无对应配置", c)
	}

}
