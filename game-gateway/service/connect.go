package service

import (
	"fmt"
	"game-gateway/manager"
	"game-gateway/protoc/pb"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"net/http"
)

func (m *gameClientManager) RoomRegister(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	if len(userId) <= 0 {
		fmt.Println("缺少userId")
		//return
	}
	///userId = utils.GetUUid()
	if len(userId) == 0 {
		userId = "syn"
	}
	upgrader := websocket.Upgrader{
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		fmt.Println("Upgrade err :", err)
		return
	}

	gameClient := NewGameClient(conn, userId)

	m.register <- gameClient

	go gameClient.Read()
	go gameClient.Write()

	//客户端 sub 订阅服务端 发过来的消息 再发给端侧
	//manager.NastManager.SubTopic("topic-syn")
	manager.NastManager.Nats.Subscribe("topic-syn", func(msg *nats.Msg) {
		netMsg := &pb.NetMessage{}
		err := proto.Unmarshal(msg.Data, netMsg)
		if err != nil {
			fmt.Println("SubTopic err ", err)
		}
		fmt.Println("接收到 服务端 发来的信息 netMessage == ", netMsg)

		// 找到userid
		//发 往 conn manager sendOut
		//然后 长链接 管理器 找到 user对应的 连接
		m.SendOutBack(netMsg)

	})

}

func (m *gameClientManager) SendOutBack(msg *pb.NetMessage) {
	m.sentOut <- msg
}
