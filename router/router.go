package router

import (
	"github.com/duiying/go-demo/module/agent"
	"github.com/duiying/go-demo/module/test"
	"github.com/duiying/go-demo/module/user"
	"github.com/duiying/go-demo/pkg/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(app *gin.Engine) *gin.Engine {
	app.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello Gin")
	})

	// 404
	app.NoRoute(middleware.NotFound())

	// pprof
	pprof.Register(app)

	// 用户相关
	app.GET("/user/find", user.Find)
	app.GET("/user/search", user.Search)
	app.POST("/user/update", user.Update)
	app.POST("/user/create", user.Create)
	app.GET("/user/redis", user.Redis)

	// 测试
	app.GET("/test", test.CustomTest)

	// WebSocket
	agent.InitHub()
	go agent.RunHub()
	app.GET("/agent", agent.WS)

	return app
}
