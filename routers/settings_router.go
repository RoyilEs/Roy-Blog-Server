package routers

import (
	"Goblog/api"
)

func (router Group) SettingsRouter() {
	settingApi := api.ApiGroupApp.SettingApi
	router.GET("settings/:name", settingApi.SettingsInfoView)
	router.PUT("settings", settingApi.SettingsInfoUpdateView)
	router.PUT("settings_qiniu", settingApi.SettingsQiNiuUpdateView)
	router.PUT("settings_jwt", settingApi.SettingsJwtUpdateView)
	router.PUT("settings_admin", settingApi.SettingsAdminUpdateView)
}
