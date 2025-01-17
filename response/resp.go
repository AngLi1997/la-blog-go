package response

import "github.com/gin-gonic/gin"

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	successCode      = 200
	failCode         = -1
	errorCode        = 500
	unauthorizedCode = 401
)

func Success(context *gin.Context, msg string) {
	context.JSON(200, Resp{successCode, msg, nil})
}

func SuccessWithData(context *gin.Context, msg string, data interface{}) {
	context.JSON(200, Resp{successCode, msg, data})
}

func Fail(context *gin.Context, msg string) {
	context.JSON(200, Resp{failCode, msg, nil})
}

func FailWithData(context *gin.Context, msg string, data interface{}) {
	context.JSON(200, Resp{failCode, msg, data})
}

func Error(context *gin.Context, msg string) {
	context.JSON(200, Resp{errorCode, msg, nil})
}

func Unauthorized(context *gin.Context) {
	context.JSON(200, Resp{unauthorizedCode, "未认证", nil})
}
