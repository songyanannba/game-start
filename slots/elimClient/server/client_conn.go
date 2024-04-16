package server

import (
	"elimClient/pbs/common"
	"elimClient/util"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type gameClient struct {
	ID         string
	UserId     string
	SocketConn *websocket.Conn
	Send       chan []byte
	ReadChan   chan *common.NetMessage
	SendText   chan []byte
	CloseChan  chan struct{}
	Closeted   bool
}

func NewGameClient(conn *websocket.Conn, userId string) *gameClient {
	return &gameClient{
		ID:         util.GetUUid(),
		UserId:     userId,
		SocketConn: conn,
		Send:       make(chan []byte, 102400),
		ReadChan:   make(chan *common.NetMessage, 102400),
		SendText:   make(chan []byte, 102400),
		CloseChan:  make(chan struct{}),
		Closeted:   false,
	}
}

func (c *gameClient) Read() {
	defer func() {
		c.SocketConn.Close()
	}()

	for {
		mtype, msg, err := c.SocketConn.ReadMessage()
		if err != nil {
			fmt.Println("c.SocketConn.ReadMessage err", err)
			break
		}
		if mtype == websocket.BinaryMessage {
			message := &common.NetMessage{}
			err := proto.Unmarshal(msg, message)
			if err == nil {
				fmt.Println("接收到协议", message.Type)
			} else {
				fmt.Println("gameClient Read proto.Unmarshal err")
				break
			}

			//manager.NastManager.Send(SERVICE, message)
			c.ReadChan <- message

		} else if mtype == websocket.PingMessage {
			c.SocketConn.PongHandler()
		} else {
			fmt.Println("read socket mtype ", mtype)
			break
		}
	}
}

func (c *gameClient) Write() {

	defer func() {
		c.SocketConn.Close()
	}()

	//收到服务返回的信息 并返回
	for {
		select {
		case data, ok := <-c.Send:
			if ok {
				fmt.Println("接收客户端管理发过来的 数据", string(data))
				c.SocketConn.WriteMessage(websocket.BinaryMessage, data)
			} else {
				fmt.Println("gameClient Write <-c.Send  err")
			}
		}
	}

}
