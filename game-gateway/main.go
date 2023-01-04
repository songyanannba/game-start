package main

import "game-gateway/api"

func main() {

	initMain()

	newApi := api.NewApi()

	newApi.Run()
}

func initMain() {
	//配置

	//db

	//日志

	//...
}
