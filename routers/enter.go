package routers

import (
	"Goblog/global"
	"github.com/gin-gonic/gin"
)

type Group struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiGroup := router.Group("api")
	routerGroup := Group{apiGroup}
	// 系统配置api
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()
	return router
}
