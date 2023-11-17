package images_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"Goblog/utils"
	"github.com/gin-gonic/gin"
)

type ImageAllResponse struct {
	Pid uint   `json:"pid"`
	Url string `json:"url"`
}

func (ImageApi) ImageAllGetOneView(c *gin.Context) {
	var imageAllModel models.ImageAll
	r := utils.Random(0, 127)
	err := global.DB.Where("id = ?", uint(r)).Find(&imageAllModel).Error
	if err != nil {
		res.ResultFailWithMsg("查询失败", c)
		return
	}
	imageAllResponse := ImageAllResponse{
		Pid: imageAllModel.Pid,
		Url: imageAllModel.Url,
	}
	res.ResultOkWithData(imageAllResponse, c)
}
