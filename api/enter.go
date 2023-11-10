package api

import (
	"Goblog/api/article_api"
	"Goblog/api/images_api"
	"Goblog/api/menu_api"
	"Goblog/api/settings_api"
)

// ApiGroup Api组件 全部实例
type ApiGroup struct {
	SettingApi settings_api.SettingsApi
	ImagesApi  images_api.ImageApi
	ArticleApi article_api.ArticleApi
	MenuApi    menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
