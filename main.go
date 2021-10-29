package main

import (
	"github.com/duiying/go-demo/router"
)

func main() {
	router := router.SetupRouter()
    _ = router.Run()
}