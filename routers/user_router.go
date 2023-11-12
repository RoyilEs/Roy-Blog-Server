package routers

import "Goblog/api"

func (router Group) UserRouter() {
	UserApi := api.ApiGroupApp.UserApi
	router.POST("email_login", UserApi.EmailLoginView)
}
