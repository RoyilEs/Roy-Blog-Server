package images_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件ID"`
	Name string `json:"name" binding:"required" msg:"请输入文件名"`
}

func (ImageApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
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
}
func ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
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
}
