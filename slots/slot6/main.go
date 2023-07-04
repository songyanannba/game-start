package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"slot6/protoc/pb"
	"slot6/src"
	"syscall"
)

func main() {

	grpcHold1()

}

// grpc 服务
func grpcHold1() {
	server := grpc.NewServer()
	pb.RegisterSlotServiceServer(server, &src.SlotService{})

	addr := fmt.Sprintf("%s:%d", "192.168.6.120", 9001)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("grpc err")
		panic(err)
	}
	fmt.Println("grpc suc...")

	err = server.Serve(listen)
	if err != nil {
		fmt.Println("grpc err...")
		panic(err)
	}
}

// 常规保持服务不退出
func hold1() {
	controlC := make(chan os.Signal, 1)
	signal.Notify(controlC, os.Interrupt, syscall.SIGTERM)
	service := src.SlotService{}
	service.Start()
	fmt.Println("slot6 启动成功。。。")
	for {
		select {
		case <-controlC:
			fmt.Println("slot6 end")
			return
		}
	}
}
