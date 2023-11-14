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
	router.PUT("images", imagesApi.ImageUpdateView)
}
