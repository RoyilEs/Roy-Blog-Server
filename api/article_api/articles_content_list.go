package article_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type ContentListResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Source   string `json:"source"`
	Word     string `json:"word"`
	Nickname string `json:"nickname"`
}

func (ArticleApi) ArticleContentListView(c *gin.Context) {
	var ContentList []ContentListResponse
	global.DB.Model(models.ArticleModel{}).Select("id", "title", "abstract", "content", "category", "source", "word", "nickname").Scan(&ContentList)
	res.ResultOkWithData(ContentList, c)
}
