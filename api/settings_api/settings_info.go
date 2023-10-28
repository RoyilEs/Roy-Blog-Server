package settings_api

import (
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsInfoView 视图
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.ResultFailWithCode(2, c)
}
