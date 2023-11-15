package routers

import (
	"Goblog/api"
	"Goblog/middleware"
)

func (router Group) ArticlesRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("articles", articleApi.ArticleListView)
	router.GET("article_content_list", articleApi.ArticleContentListView)
	router.PUT("articles/:id", articleApi.ArticleUpdateView)
	router.DELETE("articles", articleApi.ArticleRemoveView)
}
