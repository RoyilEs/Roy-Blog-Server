package routers

import "Goblog/api"

func (router Group) ArticlesRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("articles", articleApi.ArticleUploadView)
}
