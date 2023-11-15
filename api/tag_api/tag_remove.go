package tag_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}

	var TagList []models.Tag
	count := global.DB.Find(&TagList, cr.IDList).RowsAffected
	if count == 0 {
		res.ResultFailWithMsg("标签不存在", c)
		return
	}
	global.DB.Delete(&TagList)
	res.ResultOkWithMsg(fmt.Sprintf("共删除 %d 个标签", count), c)
}
