package menu_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	//更新
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.ResultFailWithMsg("菜单不存在", c)
		return
	}
	//结构体转map 第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("修改菜单失败", c)
		return
	}
	res.ResultOkWithMsg("修改菜单成功", c)
}
