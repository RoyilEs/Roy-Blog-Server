package routers

import (
	"Goblog/api"
	"Goblog/middleware"
)

func (router Group) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("email_login", userApi.EmailLoginView)
	router.POST("users", userApi.UserCreateView)
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("user_update_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("user_update_password", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
	router.POST("logout", middleware.JwtAuth(), userApi.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), userApi.UserRemoveView)

}
