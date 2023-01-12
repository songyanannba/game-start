package main

import (
	"fmt"
	"game-gateway/api"
	"game-gateway/manager"
	"game-gateway/service"
)

func main() {

	fmt.Println("start")

	initMain()

	newApi := api.NewApi()

	fmt.Println("123435657890")
	newApi.Run("127.0.0.1:8765")
}

func initMain() {
	//配置

	//db

	//日志

	//...
	manager.NastManager.Start()

	service.GameClientManager.Start()
}
