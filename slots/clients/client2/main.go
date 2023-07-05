package main

import (
	"client2/service"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	controlC := make(chan os.Signal, 1)
	signal.Notify(controlC, os.Interrupt, syscall.SIGTERM)

	service.CommonService.Start()
	service.CliHandler.Start()

	service.WsClientService.Start()

	fmt.Println("启动成功。。。")
	for {
		select {
		case <-controlC:
			return
		}
	}

}
