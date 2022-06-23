package middleware

import (
	"bytes"
	"fmt"
	"github.com/duiying/go-demo/pkg/logger"
	"github.com/duiying/go-demo/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/xiam/to"
	"runtime"
)

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

// 错误信息从 「/src/runtime/panic.go」 的下一行开始展示
var panicStart = []byte("/src/runtime/panic.go")
var lineEnd = []byte("\n")

// Recover 全局异常处理中间件
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 为了能完整输出栈信息，对 buf 进行扩充重试
				buf := make([]byte, 2048)
				for {
					n := runtime.Stack(buf, false)
					if n < len(buf) {
						buf = buf[:n]
						break
					}
					buf = make([]byte, 2*len(buf))
				}

				// 裁剪错误信息，从 panicStart 的下一行开始
				index := bytes.Index(buf, panicStart)
				if index >= 0 {
					buf = buf[index:]
					index = bytes.Index(buf, lineEnd) + 1
					buf = buf[index:]
				}

				// 记录错误堆栈信息
				logger.Error("panic", "err", fmt.Sprintf("%s\n%s", errorToString(r), to.String(buf)))

				c.JSON(500, gin.H{
					"code":    response.ServerInternalError,
					"message": response.GetMessageByCode(response.ServerInternalError),
					"data":    "",
					"traceId": c.GetString("traceId"),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
