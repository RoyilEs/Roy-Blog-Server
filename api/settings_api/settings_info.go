package settings_api

import (
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsInfoView 视图
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.ResultOkWithData(gin.H{
		"event": "settings_info",
		"姓名":    "sb 李磊",
		"我成功了":  "李磊减少十年寿命",
	}, c)
}
