package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	controlC := make(chan os.Signal, 1)
	signal.Notify(controlC)

	fmt.Println("room 启动成功。。。")
	for {
		select {
		case <-controlC:
			fmt.Println("room 结束。。。")
			return
		}
	}
}
