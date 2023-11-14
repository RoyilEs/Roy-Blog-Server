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
