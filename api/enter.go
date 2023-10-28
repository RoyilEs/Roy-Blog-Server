package api

import "Goblog/api/settings_api"

// ApiGroup Api组件 全部实例
type ApiGroup struct {
	SettingApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
