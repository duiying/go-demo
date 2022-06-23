package test

import (
	"fmt"
	"github.com/duiying/go-demo/pkg/response"
	"github.com/gin-gonic/gin"
)

func CustomTest(c *gin.Context) {
	// 测试 recover 中间件
	a := make([]int, 10)
	fmt.Println(a[10])
	response.Success(c, nil)
}
