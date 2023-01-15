package service

import (
	"fmt"
	"game-player/db"
	"game-player/protoc/pb"
	"sync"
	"time"
)

type playerSpace struct {
	PlayerInfo  *db.PlayerInfo
	CurrentTime time.Time
}

var PlayerService = playerService{
	Players:       make(map[string]*playerSpace),
	PlayerHandler: make(map[int32]*func(ps *playerSpace, msg *pb.NetMessage)),
	Sync:          new(sync.Mutex),
}

type playerService struct {
	Players       map[string]*playerSpace
	PlayerHandler map[int32]*func(ps *playerSpace, msg *pb.NetMessage)

	Sync *sync.Mutex
}

func (ps *playerService) PlayerServiceTestSend() {
	m := &pb.GameMessage{
		Name: "看绝代风华开始绝代风华空间",
	}
	NatsManager.SendGateway("syn", m)
}

func (ps *playerService) Start() {

	//注册handler
	ps.handlerInit()
}

func (ps *playerService) handlerInit() {
	ps.PlayerTest1()
}

func (ps *playerService) RegisterHandler(protoc int32, fu func(ps *playerSpace, msg *pb.NetMessage)) {
	ps.Sync.Lock()
	ps.Sync.Unlock()

	if _, ok := ps.PlayerHandler[protoc]; !ok {
		ps.PlayerHandler[protoc] = &fu
		fmt.Println("注册成功 协议 == ", protoc)
	}

}
