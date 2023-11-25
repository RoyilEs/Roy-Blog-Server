package user_api

import (
	"Goblog/models"
	"Goblog/models/ctype"
	"Goblog/models/res"
	CODE "Goblog/models/res/code"
	"Goblog/service/common"
	"Goblog/utils"
	"Goblog/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.User
	RoleID int `json:"role_id"`
}

func (UserApi) UserListView(c *gin.Context) {
	//断言
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.ResultFailWithCode(CODE.ArgumentError, c)
		return
	}
	var users []UserResponse
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
		users = append(users, UserResponse{
			User:   user,
			RoleID: int(user.Permission),
		})
	}

	res.ResultOkWithList(users, count, c)
}
