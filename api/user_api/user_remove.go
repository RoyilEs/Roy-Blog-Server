package user_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	var userList []models.User
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.ResultFailWithMsg("用户不存在", c)
		return
	}
	err = global.DB.Delete(&userList).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("用户删除失败", c)
		return
	}
	res.ResultOkWithMsg(fmt.Sprintf("成功删除%d个用户", count), c)
}
