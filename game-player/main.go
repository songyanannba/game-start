package main

import (
	"fmt"
	"game-player/service"
	"os"
	"os/signal"
)

func main() {

	controlC := make(chan os.Signal, 1)
	signal.Notify(controlC)

	service.NatsManager.Start()
	service.PlayerService.Start()

	fmt.Println("player 启动成功。。。")
	for {
		select {
		case <-controlC:
			return
		}
	}

}
