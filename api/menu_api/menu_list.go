package menu_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type MenuResponse struct {
	models.MenuModel
}

func (MenuApi) MenuListView(c *gin.Context) {
	//先查询菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)

	var menus []MenuResponse
	for _, model := range menuList {
		//model就是一个菜单
		menus = append(menus, MenuResponse{
			MenuModel: model,
		})
	}
	res.ResultOkWithData(menus, c)
	return
}
