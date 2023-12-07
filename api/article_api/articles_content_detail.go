package article_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type ArticleResponse struct {
	ContentResponse
}

func (ArticleApi) ArticleContentDetailView(c *gin.Context) {
	id := c.Param("id")
	var articleContentModel ContentResponse
	err := global.DB.Model(models.ArticleModel{}).Take(&articleContentModel, id).Error
	if err != nil {
		res.ResultFailWithMsg("文章不存在", c)
		return
	}
	articleResponse := ArticleResponse{
		ContentResponse: articleContentModel,
	}
	res.ResultOkWithData(articleResponse, c)
	return
}
