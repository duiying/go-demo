package main

import (
	"github.com/duiying/go-demo/pkg/config"
	"github.com/duiying/go-demo/pkg/middleware"
	"github.com/duiying/go-demo/pkg/mysql"
	"github.com/duiying/go-demo/pkg/redis"
	"github.com/duiying/go-demo/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("GIN_MODE") == "release" {
		// 生产模式
		gin.SetMode(gin.ReleaseMode)
		config.Debug = false
	}

	app := gin.New()

	// 全局日志中间件，使用自定义的日志中间件，可以打印响应内容
	// app.Use(gin.Logger())
	app.Use(middleware.Recover(), middleware.TraceId(), middleware.Logger())
	// 路由
	app = router.Init(app)
	// MySQL
	mysql.InitMySQL()
	// Redis
	redis.InitRedis()

	addr := ":" + os.Getenv("HTTP_PORT")
	_ = app.Run(addr)
}
