package tag_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` //显示的标题
}

func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}

	//重复的判断
	var tagModel models.Tag
	err := global.DB.Take(&tagModel, "title = ?", cr.Title).Error
	if err == nil {
		res.ResultFailWithMsg("该标签已存在", c)
		return
	}
	err = global.DB.Create(&models.Tag{Title: cr.Title}).Error
	if err != nil {
		res.ResultFailWithMsg("添加标签失败", c)
		return
	}
	res.ResultFailWithMsg("添加标签成功", c)
}
