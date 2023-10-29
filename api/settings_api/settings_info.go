package settings_api

import (
	"Goblog/global"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsInfoView 视图
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.ResultOkWithData(global.Config.SiteInfo, c)
}
