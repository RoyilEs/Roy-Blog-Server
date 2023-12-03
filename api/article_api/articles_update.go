package article_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}
	id := c.Query("id")
	fmt.Println(id)
	//更新
	var articleModel models.ArticleModel
	var bannerModel models.BannerModel
	global.DB.Take(&bannerModel, "id = ?", cr.BannerID)

	err = global.DB.Take(&articleModel, id).Error
	if err != nil {
		res.ResultFailWithMsg("文章不存在", c)
		return
	}
	//结构体转map 第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&articleModel).Updates(maps).Update("banner_path", bannerModel.Path).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("修改文章失败", c)
		return
	}
	res.ResultOkWithMsg("修改文章成功", c)
}
