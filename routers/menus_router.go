package routers

import "Goblog/api"

func (router Group) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("menus", menuApi.MenuCreateView)
	router.GET("menus", menuApi.MenuListView)
	router.GET("menus_names", menuApi.MenuNameListView)
	router.PUT("menus/:id", menuApi.MenuUpdateView)
	router.DELETE("menus", menuApi.MenuRemoveView)
}
