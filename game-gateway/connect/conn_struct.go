package connect

import (
	"golang.org/x/net/websocket"
	"sync"
)

type gameClient struct {
	ID string
	UserId string
	conn *websocket.Conn
	Send []byte
	SendText []byte
	CloseChan chan struct{}
	CloseBool bool
}

type gameClientManager struct {
	userMap map[string]map[string]*gameClient
	register chan *gameClient
	unRegister chan *gameClient
	sync sync.RWMutex
}

var GameClientManager = gameClientManager{
	userMap:    make(map[string]map[string]*gameClient),
	register:   make(chan *gameClient , 1024),
	unRegister: make(chan *gameClient , 1024),
}

func (gcm *gameClientManager) Start()  {

}