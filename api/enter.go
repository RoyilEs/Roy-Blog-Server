package api

import "Goblog/api/settings_api"

type ApiGroup struct {
	SettingApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
