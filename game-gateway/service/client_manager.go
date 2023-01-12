package service

import (
	"fmt"
	"game-gateway/protoc/pb"
	"github.com/golang/protobuf/proto"
	"sync"
)

type gameClientManager struct {
	userMap     map[string]map[string]*gameClient
	clientGroup map[string]map[string]*gameClient
	register    chan *gameClient
	unRegister  chan *gameClient
	broadCost   chan *broadCast
	sentOut     chan *pb.NetMessage
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
	userMap:     make(map[string]map[string]*gameClient),
	clientGroup: make(map[string]map[string]*gameClient),
	register:    make(chan *gameClient, 1024),
	unRegister:  make(chan *gameClient, 1024),
	sentOut:     make(chan *pb.NetMessage, 1024),
	handler:     make(map[int]func(gct *gameClient, content []byte)),
}

func (gcm *gameClientManager) Start() {

	fmt.Println("gameClientManager start...")

	for {
		select {
		case gc := <-gcm.register:
			if gcm.userMap[gc.UserId] == nil {
				gcm.userMap[gc.UserId] = make(map[string]*gameClient)
				gcm.userMap[gc.UserId][gc.ID] = gc
			} else {
				gcm.userMap[gc.UserId][gc.ID] = gc
			}

		case gc := <-gcm.unRegister:

			if uMaps, ok := gcm.userMap[gc.UserId]; ok {
				for clientId, _ := range uMaps {
					close(gcm.userMap[gc.UserId][clientId].Send)
					delete(gcm.userMap[gc.UserId], clientId)
				}
			}

		case sendData, ok := <-gcm.sentOut:
			if !ok {
				return
			}

			if mapClientIds, ok := gcm.userMap[sendData.UId]; ok {
				for id, mapClientId := range mapClientIds {
					fmt.Println("client manager send to client id", id)

					marshal, _ := proto.Marshal(sendData)
					mapClientId.Send <- marshal
				}
			}

		}

	}

}
