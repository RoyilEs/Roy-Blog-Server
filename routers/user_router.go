package routers

import (
	"Goblog/api"
	"Goblog/middleware"
)

func (router Group) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("email_login", userApi.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("user_update_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("user_update_password", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
}
