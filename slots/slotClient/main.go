package main

import (
	"fmt"

	"slotClient/server"

	"github.com/gin-gonic/gin"
	"slotClient/hander"
)

func initMain() {
	//配置
	//db
	//日志
	server.GameClientManager.Start()
}

func main() {
	initMain()

	r := gin.Default()

	accountGroup := r.Group("/v1/slot")
	{
		accountGroup.GET("/SlotTest", hander.SlotTest)      //测试grpc
		accountGroup.GET("/Slot6Spin", hander.Slot6Spin)    //短链接 slot grpc 调试
		accountGroup.GET("/Slot6SpinConn", server.GameConn) //长链接 slot grpc ; 用slot/clients/ 调试
	}

	r.Run(fmt.Sprintf("%s:%d", "127.0.0.1", 9002))
}
