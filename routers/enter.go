package routers

import (
	"Goblog/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	return router
}
