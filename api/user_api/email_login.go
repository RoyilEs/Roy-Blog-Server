package user_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"Goblog/utils/jwts"
	"Goblog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"username" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	//参数绑定 是否表中
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}

	var userModel models.User
	err = global.DB.Take(&userModel, "username = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		//没找到
		global.Log.Warn("用户名不存在")
		res.ResultFailWithMsg("用户名或密码错误", c)
		return
	}
	//校验密码
	ok := pwd.ComparePasswords(userModel.Password, cr.Password)
	if !ok {
		global.Log.Warn("用户名密码错误")
		res.ResultFailWithMsg("用户名或密码错误", c)
		return
	}

	//登录成功生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: userModel.Username,
		Nickname: userModel.Nickname,
		Role:     int(userModel.Permission),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error("token生成失败", err)
		res.ResultFailWithMsg("token生成失败", c)
		return
	}
	res.ResultOkWithData(token, c)
}
