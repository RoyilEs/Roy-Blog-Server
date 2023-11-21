package data_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	ArticleCount int64 `json:"article_count"`
	UserCount    int64 `json:"user_count"`
	ImgCount     int64 `json:"img_count"`
}

func (DataApi) DataSumView(c *gin.Context) {
	var cr DataResponse
	var (
		articleModel models.ArticleModel
		userModel    models.User
		imgModel     models.BannerModel
	)

	err := global.DB.Model(&articleModel).Count(&cr.ArticleCount).Error
	if err != nil {
		res.ResultFailWithMsg("获取文章总数失败", c)
		return
	}
	err = global.DB.Model(&userModel).Count(&cr.UserCount).Error
	if err != nil {
		res.ResultFailWithMsg("获取用户总数失败", c)
		return
	}
	err = global.DB.Model(&imgModel).Count(&cr.ImgCount).Error
	if err != nil {
		res.ResultFailWithMsg("获取图片总数失败", c)
		return
	}
	res.ResultOkWithData(cr, c)
}
