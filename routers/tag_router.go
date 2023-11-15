package routers

import "Goblog/api"

func (router Group) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	router.POST("tags", tagApi.TagCreateView)
	router.GET("tags", tagApi.TagListView)
	router.PUT("tags/:id", tagApi.TagUpdateView)
	router.DELETE("tags", tagApi.TagRemoveView)
}
