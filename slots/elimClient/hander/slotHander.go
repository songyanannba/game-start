package hander

import (
	"context"
	"elimClient/pbs/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"net/http"
)

func SlotTest(c *gin.Context) {

	//grpc 请求
	//conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	//if err != nil {
	//	fmt.Println("slit test grpc dial err", err)
	//	return
	//}
	//
	//client := pb.NewSlotServiceClient(conn)
	//client :=  common.NewElimServiceClient(conn)

	//test, err := client.SlotTest(context.Background(), &common.SlotTestReq{
	//	SID:  111,
	//	Bet:  101,
	//	Type: 3,
	//})
	//if err != nil {
	//	fmt.Println("client.SlotTest err", err)
	//	return
	//}

	//fmt.Println("client test  === ", test)

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

	client := common.NewElimServiceClient(conn)
	slot6Spin, err := client.SlotSpin(context.Background(), &common.SpinReq{
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

func Slot5Spin(c *gin.Context) {

	//grpc 请求
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("slit test grpc dial err", err)
		return
	}
	client := common.NewElimServiceClient(conn)

	slot5Spin, err := client.SlotSpin(context.Background(), &common.SpinReq{
		GameId:    5,
		SessionId: 0,
		Uid:       1,
		FreeNum:   1,
		ResNum:    1,
		Raise:     1,
		Bet:       100,
	})
	if err != nil {
		fmt.Println("client.slot5Spin err", err)
		return
	}

	var resMsg common.MatchSpinAck
	err = proto.Unmarshal(slot5Spin.MsgData.Content, &resMsg)
	if err != nil {
		fmt.Println("client Unmarshal err  === ", err)
	}
	//fmt.Println("client mmsg  === ", mmsg)
	//fmt.Println("client slot5Spin  === ", slot5Spin)

	c.JSON(http.StatusOK, gin.H{
		"mag":   "ok",
		"total": "0",
		"data2": slot5Spin.MsgData.Content,
		"data":  resMsg,
	})
}

func Slot6SpinTest() {

	//grpc 请求
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("slit test grpc dial err", err)
		return
	}

	client := common.NewElimServiceClient(conn)

	_, err = client.SlotSpin(context.Background(), &common.SpinReq{
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

	client := common.NewElimServiceClient(conn)

	slot6Spin, err := client.SlotSpin(context.Background(), &common.SpinReq{
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
