package settings_api

import (
	"Goblog/global"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsQiNiuInfoView(c *gin.Context) {
	res.ResultOkWithData(global.Config.QiNiu, c)
}
