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
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
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

func Slot6Spin(c *gin.Context) {

	//grpc 请求
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("slit test grpc dial err", err)
		return
	}

	client := pb.NewSlotServiceClient(conn)
	slot6Spin, err := client.SlotSpin(context.Background(), &pb.SpinReq{
		GameId:    6,
		SessionId: 0,
		Uid:       1,
		FreeNum:   1,
		ResNum:    1,
		Raise:     1,
		Bet:       100,
	})
	if err != nil {
		fmt.Println("client.slot6Spin err", err)
		return
	}

	fmt.Println("client slot6Spin  === ", slot6Spin)

	c.JSON(http.StatusOK, gin.H{
		"mag":   "ok",
		"total": "1001",
		"data":  "xxxxxxxxttt",
	})
}

func Slot6SpinTest() {

	//grpc 请求
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("slit test grpc dial err", err)
		return
	}

	client := pb.NewSlotServiceClient(conn)

	_, err = client.SlotSpin(context.Background(), &pb.SpinReq{
		GameId:    6,
		SessionId: 0,
		Uid:       1,
		FreeNum:   1,
		ResNum:    1,
		Raise:     1,
		Bet:       100,
	})
	if err != nil {
		fmt.Println("client.slot6Spin err", err)
		return
	}

	//fmt.Println("client slot6Spin  === ", slot6Spin)
}

func Slot6SpinConn(c *gin.Context) {
	//grpc 请求
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("slit test grpc dial err", err)
		return
	}

	client := pb.NewSlotServiceClient(conn)

	slot6Spin, err := client.SlotSpin(context.Background(), &pb.SpinReq{
		GameId:    6,
		SessionId: 0,
		Uid:       1,
		FreeNum:   1,
		ResNum:    1,
		Raise:     1,
		Bet:       100,
	})
	if err != nil {
		fmt.Println("client.slot6Spin err", err)
		return
	}

	fmt.Println("client slot6Spin  === ", slot6Spin)

	c.JSON(http.StatusOK, gin.H{
		"mag":   "ok",
		"total": "1001",
		"data":  "xxxxxxxxttt",
	})
}
