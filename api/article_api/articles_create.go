package article_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"Goblog/service/user_ser"
	"Goblog/utils/jwts"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleRequest struct {
	Title    string `json:"title"`     //文章标题
	Abstract string `json:"abstract"`  //文章介绍
	Content  string `json:"content"`   //文章内容
	Category string `json:"category"`  //文章分类
	BannerID uint   `json:"banner_id"` //封面ID
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	//断言
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr ArticleRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	//建立数据入库
	articleModel := models.ArticleModel{
		Title:    cr.Title,
		Abstract: cr.Abstract,
		Content:  cr.Content,
		Category: cr.Category,
		UserID:   claims.UserID,
		Source:   claims.Username,
		Word:     strconv.Itoa(len(cr.Content)),
		Nickname: claims.Nickname,
	}

	var (
		bannerModel models.BannerModel
		bannerPath  = user_ser.AVATAR
		bannerID    = cr.BannerID
	)
	count := global.DB.Take(&bannerModel, "id = ?", bannerID).RowsAffected
	//不存在
	if count == 0 {
		articleModel.BannerID = uint(5)
		articleModel.BannerPath = bannerPath
	} else {
		articleModel.BannerID = bannerID
		articleModel.BannerPath = bannerPath
	}

	err := global.DB.Create(&articleModel).Error
	if err != nil {
		res.ResultFailWithMsg("文章添加失败", c)
		return
	}
	res.ResultOkWithMsg("文章添加成功", c)
}
