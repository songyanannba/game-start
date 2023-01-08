package service

import (
	"fmt"
	"game-gateway/manager"
	"game-gateway/protoc/pb"
	"game-gateway/util"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"sync"
)

const PLAYER = "player"
const SERVICE = "service_syn"

func NewGameClient(conn *websocket.Conn, userId string) *gameClient {
	return &gameClient{
		ID:         util.GetUUid(),
		UserId:     userId,
		SocketConn: conn,
		Send:       make([]byte, 102400),
		SendText:   make([]byte, 102400),
		CloseChan:  make(chan struct{}),
		Closeted:   false,
	}
}

type gameClient struct {
	ID         string
	UserId     string
	SocketConn *websocket.Conn
	Send       []byte
	SendText   []byte
	CloseChan  chan struct{}
	Closeted   bool
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
			message := &pb.NetMessage{}
			err := proto.Unmarshal(msg, message)

			if err == nil {
				fmt.Println("接收到协议", message.Type)
			} else {
				fmt.Println("gameClient Read proto.Unmarshal err")
				break
			}

			manager.NastManager.Send(SERVICE, message)

		} else if mtype == websocket.CloseMessage {
			//c.SocketConn.PongHandler()
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
	select {}

}

type gameClientManager struct {
	userMap     map[string]map[string]*gameClient
	clientGroup map[string]map[string]*gameClient
	register    chan *gameClient
	unRegister  chan *gameClient
	broadCost   *broadCast
	sentOut     *pb.NetMessage
	handler     map[int]func(gct *gameClient, content []byte)
	sync        sync.RWMutex
}

type broadCast struct {
	Msg     []byte
	GroupId string
	bl      bool
}

type NetMessage struct {
	Type      int
	content   []byte
	ServiceID string
	UId       string
}

var GameClientManager = gameClientManager{
	userMap:    make(map[string]map[string]*gameClient),
	register:   make(chan *gameClient, 1024),
	unRegister: make(chan *gameClient, 1024),
	handler:    make(map[int]func(gct *gameClient, content []byte)),
}

func (gcm *gameClientManager) Start() {

}
