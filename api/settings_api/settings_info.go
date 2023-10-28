package settings_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SettingsInfoView 视图
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"event": "SettingsInfoView",
	})
}
