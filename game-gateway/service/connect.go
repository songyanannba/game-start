package service

import (
	"fmt"
	"game-gateway/manager"
	"game-gateway/protoc/pb"
	"game-gateway/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
)

func RoomRegister(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	if len(userId) <= 0 {
		fmt.Println("缺少userId")
		//return
	}
	userId = util.GetUUid()

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

	go gameClient.Read()
	go gameClient.Write()

	manager.NastManager.SubTopic("topic-syn", func(msg *pb.NetMessage) {
		netMsg := &pb.NetMessage{}

		err = proto.Unmarshal(msg.Content, netMsg)
		if err != nil {
			fmt.Println("manager.NastManager.SubTopic err", err)
		}
		fmt.Println("manager.NastManager.SubTopic ", netMsg)

	})

}
