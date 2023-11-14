package user_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"Goblog/utils/jwts"
	"Goblog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"` //旧密码
	NewPassword string `json:"new_password"` //新密码
}

// UserUpdatePasswordView 修改登录人的ID
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 参数绑定
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}

	var user models.User
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.ResultFailWithMsg("用户不存在", c)
		return
	}
	// 判断密码是否一致
	if !pwd.ComparePasswords(user.Password, cr.OldPassword) {
		res.ResultOkWithMsg("密码错误", c)
		return
	}
	hashPwd := pwd.HashPassword(cr.NewPassword)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("密码修改失败", c)
		return
	}
	res.ResultOkWithMsg("密码修改成功", c)
	return
}
