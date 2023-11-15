package user_api

import (
	"Goblog/global"
	"Goblog/models/ctype"
	"Goblog/models/res"
	"Goblog/service/user_ser"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	Nickname   string     `json:"nickname" binding:"required" msg:"请输入昵称"`   //昵称
	Username   string     `json:"username" binding:"required" msg:"请输入用户名"`  //用户名
	Email      string     `json:"email"`                                     //邮箱
	Phone      string     `json:"phone"`                                     //手机号
	Password   string     `json:"password" binding:"required" msg:"请输入密码"`   //密码
	Permission ctype.Role `json:"permission" binding:"required" msg:"请选择权限"` //权限
}

func (UserApi) UserCreateView(c *gin.Context) {

	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}
	err := user_ser.UserService{}.CreateUser(cr.Username, cr.Nickname, cr.Password, cr.Permission, cr.Email, cr.Phone)
	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("用户创建失败!", c)
		return
	}

	res.ResultOkWithMsg(fmt.Sprintf("用户%s创建成功!", cr.Username), c)
	return
}
