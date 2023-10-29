package routers

import (
	"Goblog/api"
)

func (router Group) SettingsRouter() {
	settingApi := api.ApiGroupApp.SettingApi
	router.GET("settings", settingApi.SettingsInfoView)
}
