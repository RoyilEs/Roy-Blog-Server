package images_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImageApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.ResultOkWithData(imageList, c)
}
