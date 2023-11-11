package menu_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	//先查询菜单
	id := c.Param("id")
	var menuModel models.MenuModel

	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.ResultFailWithMsg("菜单不存在", c)
		return
	}
	//model就是一个菜单
	menuResponse := MenuResponse{
		MenuModel: menuModel,
	}

	res.ResultOkWithData(menuResponse, c)
	return
}
