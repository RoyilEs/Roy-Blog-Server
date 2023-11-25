package user_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/ctype"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Nickname string     `json:"nickname" binding:"required" msg:"请输入昵称"`
	Role     ctype.Role `json:"permission" binding:"required,oneof= 1 2 3 4" msg:"权限参数错误"`
	UserID   uint       `json:"user_id" binding:"required" msg:"用户ID错误"`
}

// UserUpdateRoleView 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}
	var user models.User
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.ResultFailWithMsg("用户ID错误,用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Update("Permission", cr.Role).Update("Nickname", cr.Nickname).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("修改权限错误", c)
		return
	}
	res.ResultOkWithMsg("修改成功", c)
}
