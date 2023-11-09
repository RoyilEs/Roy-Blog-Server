package article_api

import (
	"Goblog/models"
	"Goblog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleUploadView(c *gin.Context) {
	var article models.ArticleModel
	err := c.BindJSON(&article)
	fmt.Println("11111")
	if err != nil {
		res.ResultFailWithMsg(err.Error(), c)
		return
	}
	res.ResultOkWithData(article, c)
}
