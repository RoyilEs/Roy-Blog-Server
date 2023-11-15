package routers

import (
	"Goblog/api"
	"Goblog/middleware"
)

func (router Group) ArticlesRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("article_list", articleApi.ArticleListView)
}
