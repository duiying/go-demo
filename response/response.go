package response

import "github.com/gin-gonic/gin"

const ServerError = 4000
const ParamsError = 4001
const ExistError = 4002

// 错误码 & 错误信息映射
var codeMap = map[int]string{
	ServerError: "服务异常",
	ParamsError: "参数错误",
	ExistError: "记录不存在",
}

// 根据 code 返回错误信息
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
	})
	return
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
	})
	return
}

