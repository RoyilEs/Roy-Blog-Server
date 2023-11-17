package images_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"Goblog/utils"
	"github.com/gin-gonic/gin"
)

type AnosuResponse struct {
	Url string `json:"url"`
}

// AnosuOneVie TODO 表出问题
func (ImageApi) AnosuOneVie(c *gin.Context) {

	var anosuModel models.AnosuAll
	r := utils.Random(162, 3679)
	err := global.DB.Where("id = ?", r).Find(&anosuModel).Error
	if err != nil {
		res.ResultFailWithMsg("查询失败", c)
		return
	}
	anosuResponse := AnosuResponse{
		Url: anosuModel.Url,
	}
	res.ResultOkWithData(anosuResponse, c)
}
