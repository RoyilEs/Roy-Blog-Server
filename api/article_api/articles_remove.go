package article_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	var articleList []models.ArticleModel
	count := global.DB.Find(&articleList, cr.IDList).RowsAffected
	if count == 0 {
		res.ResultFailWithMsg("文章不存在", c)
		return
	}
	err = global.DB.Delete(&articleList).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("文章删除失败", c)
		return
	}
	res.ResultOkWithMsg(fmt.Sprintf("成功删除%d个文章", count), c)
}
