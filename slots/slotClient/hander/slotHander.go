package hander

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"slotClient/proto/pb"
)

func SlotTest(c *gin.Context) {

	//grpc 请求
	conn, err := grpc.Dial("192.168.6.120:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("slit test grpc dial err", err)
		return
	}

	client := pb.NewSlotServiceClient(conn)

	test, err := client.SlotTest(context.Background(), &pb.SlotTestReq{
		SID:  111,
		Bet:  101,
		Type: 3,
	})
	if err != nil {
		fmt.Println("client.SlotTest err", err)
		return
	}

	fmt.Println("client test  === ", test)

	c.JSON(http.StatusOK, gin.H{
		"mag":   "ok",
		"total": "1001",
		"data":  "xxxxxxxxttt",
	})
}
