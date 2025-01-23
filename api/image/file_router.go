package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"la-blog-go/api"
	"la-blog-go/response"
)

var (
	Apis = []api.Api{
		{
			Method: "POST",
			Url:    "/upload",
			Func: func(c *gin.Context) {
				file, _ := c.FormFile("file")
				fmt.Println(file.Filename)
				response.SuccessWithData(c, "上传成功", "http://172.30.1.160/front-end/assets/logo/bmos/Bmos_logo.svg")
			},
		},
	}
)
