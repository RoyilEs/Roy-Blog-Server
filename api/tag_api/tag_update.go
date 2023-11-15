package tag_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")

	var cr TagRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}

	var tagModel models.Tag
	err := global.DB.Take(&tagModel, id).Error
	if err != nil {
		res.ResultFailWithMsg("标签不存在", c)
		return
	}

	//结构体转map 第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&tagModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("修改标签失败", c)
		return
	}
	res.ResultOkWithMsg("修改标签成功", c)

}
