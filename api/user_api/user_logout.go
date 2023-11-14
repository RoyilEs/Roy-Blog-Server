package user_api

import (
	"Goblog/global"
	"Goblog/models/res"
	"Goblog/service"
	"Goblog/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")

	err := service.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		res.ResultFailWithMsg("注销失败", c)
		return
	}

	res.ResultOkWithMsg("注销成功", c)
}

//注销 记一次笔记 写入server 作为全局
//其中调用 redis_ser 的 logout针对注销的操作存入redis
//在中间件中判断是否登录 未登录则返回错误 获得redis中 keys 在utils.InList 判断是否在redis中 --> redis_ser.CheckLogout
//存在则为注销
