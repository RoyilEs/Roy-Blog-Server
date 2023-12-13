package menu_api

import (
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"Goblog/service/common"
	"github.com/gin-gonic/gin"
)

type MenuResponse struct {
	models.MenuModel
}

func (MenuApi) MenuListView(c *gin.Context) {

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}

	menuList, count, err := common.ComList(&models.MenuModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	if err != nil {
		return
	}
	res.ResultOkWithList(menuList, count, c)
	return
}
