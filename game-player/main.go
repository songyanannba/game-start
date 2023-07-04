package main

import (
	"fmt"
	"game-player/conf"
	"game-player/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	controlC := make(chan os.Signal, 1)
	signal.Notify(controlC, os.Interrupt, syscall.SIGTERM)

	conf.PlayerConfInit() //配置

	service.NatsManager.Start()

	service.PlayerService.Start()

	fmt.Println("player 启动成功。。。")
	for {
		select {
		case <-controlC:
			fmt.Println("player end。。。")
			return
		}
	}

}
