package user_api

import (
	"Goblog/models"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"Goblog/service/common"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.User{}, common.Option{
		PageInfo: page,
	})
	res.ResultOkWithList(list, count, c)

}
