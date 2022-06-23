package response

import "github.com/gin-gonic/gin"

const ErrorCode = -1

const ServerInternalError = 500
const ServerError = 4000
const ParamsError = 4001
const ExistError = 4002
const CreateError = 4003
const NotWebSocketError = 4004

// 错误码 & 错误信息映射
var codeMap = map[int]string{
	ServerInternalError: "内部错误",
	ServerError:         "服务异常",
	ParamsError:         "参数错误",
	ExistError:          "记录不存在",
	CreateError:         "创建失败",
	NotWebSocketError:   "not websocket",
}

// List 列表返回的结构体
type List struct {
	P     int         `json:"p"`
	Size  int         `json:"size"`
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

// GetMessageByCode 根据 code 返回错误信息
func GetMessageByCode(code int) string {
	if message, ok := codeMap[code]; ok {
		return message
	}
	return "未知错误码"
}

func Fail(c *gin.Context, code int) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": GetMessageByCode(code),
		"data":    "",
		"traceId": c.GetString("traceId"),
	})
	return
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
		"traceId": c.GetString("traceId"),
	})
	return
}
