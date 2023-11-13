package routers

import "Goblog/api"

func (router Group) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("email_login", userApi.EmailLoginView)
	router.GET("users", userApi.UserListView)
}
