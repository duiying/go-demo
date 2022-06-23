package server

import (
	"fmt"
	"github.com/duiying/go-demo/pkg/config"
	"github.com/duiying/go-demo/pkg/middleware"
	"github.com/duiying/go-demo/pkg/mysql"
	"github.com/duiying/go-demo/pkg/redis"
	"github.com/duiying/go-demo/router"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
)

var (
	port   string
	mode   string
	Server = &cobra.Command{
		Use: "server",
		PreRun: func(cmd *cobra.Command, args []string) {
			welcome()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	Server.PersistentFlags().StringVarP(&port, "port", "p", "9551", "http port server listening on")
	Server.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode; eg:dev,test,prod")
}

func welcome() {
	log.Println(color.GreenString("starting http server"))
}

func run() error {
	// 初始化 env
	err := godotenv.Load(fmt.Sprintf(".env.%s", mode))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 生产模式
	if mode == "prod" {
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

	addr := ":" + port
	_ = app.Run(addr)

	return nil
}
