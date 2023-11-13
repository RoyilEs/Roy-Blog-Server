package routers

import (
	"Goblog/api"
)

func (router Group) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.GET("images_names", imagesApi.ImageNameListView)
	router.DELETE("images", imagesApi.ImageRemoveView)
	//TODO 不知道哪里出的bug 包传递也不行了 ---- 直接匿名函数
	router.PUT("images", imagesApi.ImageUpdateView)
}
