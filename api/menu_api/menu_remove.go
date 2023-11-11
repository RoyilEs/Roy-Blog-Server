package menu_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.ResultFailWithMsg("菜单不存在", c)
		return
	}
	err = global.DB.Delete(&menuList).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("菜单删除失败", c)
		return
	}
	res.ResultOkWithMsg(fmt.Sprintf("成功删除%d个菜单", count), c)
}
