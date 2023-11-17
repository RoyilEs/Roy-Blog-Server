package routers

import (
	"Goblog/api"
)

func (router Group) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.GET("images_names", imagesApi.ImageNameListView)
	router.GET("image_get_https", imagesApi.ImageGetHttpsView)
	router.GET("anosu_one", imagesApi.AnosuOneVie)
	router.GET("imageAll_one", imagesApi.ImageAllGetOneView)
	router.DELETE("images", imagesApi.ImageRemoveView)
	router.PUT("images", imagesApi.ImageUpdateView)
}
