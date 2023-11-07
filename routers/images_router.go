package routers

import (
	"Goblog/api"
	"Goblog/api/images_api"
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

func (router Group) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", imagesApi.ImageUploadView)
	router.GET("images", imagesApi.ImageListView)
	router.DELETE("images", imagesApi.ImageRemoveView)
	//TODO 不知道哪里出的bug 包传递也不行了 ---- 直接匿名函数
	router.PUT("images", func(c *gin.Context) {
		var cr images_api.ImageUpdateRequest
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.ResultFailWithError(err, &cr, c)
			return
		}
		//拿到ID 进行修改
		var imageModel models.BannerModel
		err = global.DB.Take(&imageModel, cr.ID).Error
		if err != nil {
			res.ResultFailWithMsg("文件不存在", c)
			return
		}
		//修改
		err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
		if err != nil {
			res.ResultFailWithMsg(err.Error(), c)
			return
		}
		res.ResultOkWithMsg("文件名修改成功", c)
	})
}
