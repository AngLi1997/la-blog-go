package router

import (
	"github.com/gin-gonic/gin"
	"la-blog-go/api"
	"la-blog-go/api/article"
	"la-blog-go/api/category"
	"la-blog-go/api/tag"
	"la-blog-go/response"
)

const globalRequestPrefix = "/api"

func InitRouters() {
	r := gin.Default()
	//r.Use(TokenAuth())
	handleGroup("article", r, article.Apis...)
	handleGroup("tag", r, tag.Apis...)
	handleGroup("category", r, category.Apis...)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}

func handleGroup(group string, engine *gin.Engine, apis ...api.Api) {
	g := engine.Group(globalRequestPrefix + "/" + group)
	for index := range apis {
		g.Handle(apis[index].Method, apis[index].Url, bindParams(apis[index].Dst, apis[index].ParamType, apis[index].Func))
	}
}

func bindParams(paramDst interface{}, paramType string, afterBind func(context *gin.Context, obj interface{})) gin.HandlerFunc {
	return func(context *gin.Context) {
		if paramDst != nil {
			if paramType == "body" {
				if err := context.ShouldBindJSON(&paramDst); err != nil {
					response.Fail(context, "参数格式错误")
					return
				}
			} else if paramType == "param" {
				if err := context.ShouldBindQuery(&paramDst); err != nil {
					response.Fail(context, "参数格式错误")
					return
				}
			}
		}
		afterBind(context, paramDst)
	}
}
