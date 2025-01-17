package api

import "github.com/gin-gonic/gin"

type Api struct {
	Url    string
	Method string
	Dst    interface{}
	Func   func(c *gin.Context, dst interface{})
}
