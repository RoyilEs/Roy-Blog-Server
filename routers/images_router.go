package routers

import (
	"Goblog/api"
	"Goblog/middleware"
)

func (router Group) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.GET("images_names", imagesApi.ImageNameListView)
	router.GET("image_get_https", imagesApi.ImageGetHttpsView)
	router.GET("anosu_one", imagesApi.AnosuOneVie)
	router.GET("imageAll_one", imagesApi.ImageAllGetOneView)
	router.GET("images_one", imagesApi.ImageGetOneView)
	router.DELETE("images", middleware.JwtAdmin(), imagesApi.ImageRemoveView)
	router.PUT("images", middleware.JwtAuth(), imagesApi.ImageUpdateView)
}
