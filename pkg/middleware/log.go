package middleware

import (
	"bytes"
	"fmt"
	"github.com/duiying/go-demo/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type reqLog struct {
	Method   string
	Path     string
	ClientIP string
	TraceId  string
}

type respLog struct {
	Method     string
	Path       string
	ClientIP   string
	TraceId    string
	Cost       string
	StatusCode int
	Resp       string
}

// 格式化时间
func getFormatTime(beginTime time.Time, endTime time.Time) string {
	beginTimeMicro := beginTime.UnixMicro()
	endTimeMicro := endTime.UnixMicro()
	subTime := endTimeMicro - beginTimeMicro
	if subTime < 1000 {
		return fmt.Sprintf("%dus", subTime)
	} else {
		return fmt.Sprintf("%dms", subTime/1000)
	}
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Logger 全局日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginTime := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// 记录相应内容
		writer := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// 请求方式
		method := c.Request.Method

		traceId := c.GetString("traceId")

		req := reqLog{
			Method:   method,
			Path:     path,
			ClientIP: c.ClientIP(),
			TraceId:  traceId,
		}

		logger.Debug("request", "req", req)

		c.Next()

		endTime := time.Now()

		resp := respLog{
			Method:     method,
			Path:       path,
			ClientIP:   c.ClientIP(),
			TraceId:    traceId,
			Cost:       getFormatTime(beginTime, endTime),
			StatusCode: c.Writer.Status(),
			Resp:       writer.body.String(),
		}

		logger.Debug("response", "resp", resp)
	}
}
