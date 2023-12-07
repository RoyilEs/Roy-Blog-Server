package article_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type ContentResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Abstract   string `json:"abstract"`
	Content    string `json:"content"`
	Category   string `json:"category"`
	Source     string `json:"source"`
	Word       string `json:"word"`
	Nickname   string `json:"nickname"`
	BannerPath string `json:"banner_path"`
}

func (ArticleApi) ArticleContentListView(c *gin.Context) {
	var ContentList []ContentResponse
	global.DB.Model(models.ArticleModel{}).Select("id", "title", "abstract", "content", "category", "source", "word", "nickname", "banner_path").Scan(&ContentList)
	res.ResultOkWithData(ContentList, c)
}
