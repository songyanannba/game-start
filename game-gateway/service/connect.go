package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func RoomRegister(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	if len(userId) <= 0 {
		fmt.Println("缺少userId")
		return
	}


	upgrader := websocket.Upgrader{
		Subprotocols:      []string{ctx.GetHeader("Sec-WebSocket-Protocol")},

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade()

	if err != nil {
		fmt.Println("Upgrade err :" ,err)
		return
	}
	//up := websocket.

}
