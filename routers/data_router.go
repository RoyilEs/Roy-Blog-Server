package routers

import "Goblog/api"

func (router Group) DataRouter() {
	dataRouter := api.ApiGroupApp.DataApi
	router.GET("data_sum", dataRouter.DataSumView)
}
