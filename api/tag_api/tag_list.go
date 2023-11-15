package tag_api

import (
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"Goblog/service/common"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.Tag{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.ResultOkWithList(list, count, c)
}
