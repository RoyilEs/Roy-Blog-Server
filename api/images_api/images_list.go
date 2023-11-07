package images_api

import (
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"Goblog/service/common"
	"github.com/gin-gonic/gin"
)

// ImageListView 图片列表
// @Tags  图片
// @Summary 图片列表
// @Description 图片列表
// @Param data query model.PageInfo true "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImageApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	//var imageList []models.BannerModel
	////查询所有
	//count := global.DB.Debug().Select("id").Find(&imageList).RowsAffected
	////分页的小算法 对应 limit
	//offset := cr.Limit * (cr.Page - 1)
	//if offset < 0 {
	//	offset = 0
	//}
	//global.DB.Limit(cr.Limit).Offset(offset).Find(&imageList)
	imageList, count, err := common.ComList(&models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})

	res.ResultOkWithList(imageList, count, c)
	return
}
