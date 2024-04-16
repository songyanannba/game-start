package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"net/http"
	"slotClient/proto/pb"
)

func (m *gameClientManager) Register(ctx *gin.Context) {

	userId := ctx.GetHeader("userId") // todo 从token解析 userId
	if len(userId) <= 0 {
		fmt.Println("缺少userId")
		//return
	}
	///userId = utils.GetUUid()
	if len(userId) == 0 {
		userId = "syn"
	}

	//gameClient := CreatClient(ctx, userId) //生成 websocket 客户端

	upGrader := websocket.Upgrader{
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Upgrade err :", err)
		return
	}
	gameClient := NewGameClient(conn, userId)

	m.register <- gameClient
	go gameClient.Read()
	go gameClient.Write()

	//客户端 等待 服务端 发过来的消息 再发给端侧
	go func() {
		for {
			select {
			case data, ok := <-gameClient.ReadChan:
				if ok {
					fmt.Println("接收客户端管理发过来的 数据", data)
					//go m.DealReadChan(data)

					if data.Type == -1 { //退出
						m.unRegister <- gameClient
					} else {
						m.DealReadChan(data)
					}

				} else {
					fmt.Println("gameClient Write <-c.Send  err")
				}
			}
		}
	}()

}

func (m *gameClientManager) DealReadChan(SpinReqData *pb.NetMessage) {

	//grpc 请求 logicPack
	conn, err := grpc.Dial("192.168.6.119:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println("slit test grpc dial err", err)
		return
	}

	client := pb.NewSlotServiceClient(conn)
	slot6Spin, err := client.SlotSpin(context.Background(), &pb.SpinReq{
		GameId:    SpinReqData.SlotID,
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
	fmt.Println("SendOutBackStart slot6Spin === ", slot6Spin)

	netMsg := &pb.NetMessage{
		UId:     "syn",
		Content: []byte{1, 2, 3, 4},
		Type:    1,
	}
	//err := proto.Unmarshal(msg.Data, netMsg)
	//if err != nil {
	//	fmt.Println("SubTopic err ", err)
	//}
	fmt.Println("接收到 服务端 发来的信息 netMessage == ", netMsg)

	// 找到userid
	//发 往 conn manager sendOut
	//然后 长链接 管理器 找到 user对应的 连接
	m.SendOutBack(netMsg)

}

// CreatClient 生成 websocket 客户端
func CreatClient(ctx *gin.Context, userId string) *gameClient {
	upGrader := websocket.Upgrader{
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Upgrade err :", err)
		return nil
	}
	return NewGameClient(conn, userId)
}

func (m *gameClientManager) SendOutBack(msg *pb.NetMessage) {
	m.sentOut <- msg
}
