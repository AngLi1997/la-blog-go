package router

import (
	"github.com/gin-gonic/gin"
	"la-blog-go/api"
	"la-blog-go/api/article"
	"la-blog-go/api/category"
	"la-blog-go/api/image"
	"la-blog-go/api/tag"
)

const globalRequestPrefix = "/api"

func InitRouters() {
	r := gin.Default()
	//r.Use(TokenAuth())
	handleGroup("article", r, article.Apis...)
	handleGroup("tag", r, tag.Apis...)
	handleGroup("category", r, category.Apis...)
	handleGroup("file", r, image.Apis...)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}

func handleGroup(group string, engine *gin.Engine, apis ...api.Api) {
	g := engine.Group(globalRequestPrefix + "/" + group)
	for index := range apis {
		g.Handle(apis[index].Method, apis[index].Url, preFilter(apis[index].Func))
	}
}

func preFilter(afterBind func(context *gin.Context)) gin.HandlerFunc {
	return func(context *gin.Context) {
		afterBind(context)
	}
}
