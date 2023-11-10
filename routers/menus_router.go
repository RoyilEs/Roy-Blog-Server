package routers

import "Goblog/api"

func (router Group) MenusRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("menus", menuApi.MenuCreateView)
}
