package router

import (
	"github.com/gin-gonic/gin"
	"la-blog-go/api"
	"la-blog-go/api/article"
	"la-blog-go/api/tag"
	"la-blog-go/response"
)

func InitRouters() {
	r := gin.Default()
	//r.Use(TokenAuth())

	handleGroup("article", r, article.Apis...)
	handleGroup("tag", r, tag.Apis...)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}

func handleGroup(group string, engine *gin.Engine, apis ...api.Api) {
	g := engine.Group(group)
	for index := range apis {
		g.Handle(apis[index].Method, apis[index].Url, bindParams(apis[index].Dst, apis[index].Func))
	}
}

func bindParams(paramDst interface{}, afterBind func(context *gin.Context, obj interface{})) gin.HandlerFunc {
	return func(context *gin.Context) {
		if paramDst != nil {
			if err := context.ShouldBindJSON(&paramDst); err != nil {
				response.Fail(context, "参数错误")
				return
			}
		}
		afterBind(context, paramDst)
	}
}
