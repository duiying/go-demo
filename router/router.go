package router

import (
	"github.com/duiying/go-demo/module/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(app *gin.Engine) *gin.Engine {
	app.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})

	// 用户相关
	app.GET("/user/find", user.Find)

	return app
}