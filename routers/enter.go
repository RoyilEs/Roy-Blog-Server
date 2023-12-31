package routers

import (
	"Goblog/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

type Group struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//设置静态路由
	router.StaticFS("uploads", http.Dir("uploads"))
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("api")
	routerGroup := Group{apiGroup}
	// 系统配置api
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()
	routerGroup.ArticlesRouter()
	routerGroup.MenuRouter()
	routerGroup.UserRouter()
	routerGroup.TagRouter()
	routerGroup.DataRouter()
	return router
}
