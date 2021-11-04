package main

import (
	"github.com/duiying/go-demo/router"
	"github.com/duiying/go-demo/util"
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
	}

	app := gin.New()

	// 路由
	app = router.Init(app)
	// MySQL
	util.InitMySQL()

	addr := ":" + os.Getenv("HTTP_PORT")
	app.Run(addr)
}