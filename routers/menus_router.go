package routers

import (
	"Goblog/api"
	"Goblog/middleware"
)

func (router Group) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("menus", middleware.JwtAdmin(), menuApi.MenuCreateView)
	router.GET("menus", middleware.JwtAdmin(), menuApi.MenuListView)
	router.GET("menus_names", menuApi.MenuNameListView)
	router.PUT("menus/:id", middleware.JwtAdmin(), menuApi.MenuUpdateView)
	router.DELETE("menus", middleware.JwtAdmin(), menuApi.MenuRemoveView)
	router.GET("menus/:id", menuApi.MenuDetailView)
}
