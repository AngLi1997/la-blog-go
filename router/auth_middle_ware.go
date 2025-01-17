package router

import (
	"github.com/gin-gonic/gin"
	"la-blog-go/response"
)

func TokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			response.Unauthorized(context)
			context.Abort()
		}
		// 校验token
		context.Next()
	}
}
