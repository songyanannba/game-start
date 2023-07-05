package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"slotClient/hander"
)

func main() {
	r := gin.Default()

	accountGroup := r.Group("/v1/slot")
	{
		accountGroup.GET("/SlotTest", hander.SlotTest)
		accountGroup.GET("/Slot6Spin", hander.Slot6Spin)
	}
	r.Run(fmt.Sprintf("%s:%d", "192.168.6.119", 9002))
}
