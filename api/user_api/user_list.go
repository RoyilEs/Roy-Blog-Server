package user_api

import (
	"Goblog/models"
	"Goblog/models/ctype"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"Goblog/service/common"
	"Goblog/utils"
	"Goblog/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	//判断是否是管理员
	token := c.Request.Header.Get("token")
	if token == "" {
		res.ResultFailWithMsg("未携带token", c)
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.ResultFailWithMsg("token错误", c)
		return
	}
	fmt.Println(claims)

	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	var users []models.User
	list, count, _ := common.ComList(models.User{}, common.Option{
		PageInfo: page,
	})

	for _, user := range list {

		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//非管理员
			user.Username = ""
		}
		// 脱敏
		user.Phone = utils.DesensitizationTel(user.Phone)
		user.Email = utils.DesensitizationEmail(user.Email)
		users = append(users, user)
	}

	res.ResultOkWithList(users, count, c)
}
