package api

import "github.com/gin-gonic/gin"

type Api struct {
	Url    string
	Method string
	Func   func(c *gin.Context)
}
